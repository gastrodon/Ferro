package routes

import (
	"monke-cdn/log"
	"monke-cdn/storage"
	"monke-cdn/util"

	"crypto/md5"
	"errors"
	"mime/multipart"
	"net/http"
)

const (
	fileMemoryLimit int64 = 32 << 20
)

func getFileName(request *http.Request) (name string) {
	name = request.URL.Query().Get("name")

	if name == "" {
		// TODO do not check for UUIDv4 dupes
		name, _ = storage.NewUUID()
	}

	return
}

func getFile(request *http.Request, name string, max_size int64) (file multipart.File, mime string, err error) {
	err = request.ParseMultipartForm(max_size)

	if err != nil {
		return
	}

	var header *multipart.FileHeader
	file, header, err = request.FormFile(name)
	if header == nil {
		err = errors.New("bad_request")
		return
	}
	file, err = header.Open()
	if file == nil {
		err = errors.New("bad_request")
		return
	}

	mime = header.Header.Get("Content-Type")
	return
}

func md5File(hashable []byte) (md5_sum []byte, err error) {
	md5_sum = md5.New().Sum(hashable)
	return
}

func handleUpload(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	log.Tracef("Handling upload")

	if !util.Authed(request.Header.Get("Authorization")) {
		util.HTTPResponseError(response, "bad_auth", 401)
		return
	}

	var file multipart.File
	var mime string
	var err error
	file, mime, err = getFile(request, "file", fileMemoryLimit)
	if err != nil {
		util.HTTPResponseError(response, "bad_request", 400)
		return
	}

	var id string = getFileName(request)
	err = storage.CreateReference(id, mime)
	if err != nil {
		log.Errorf("Could not create a reference for %s (mime %s)", id, mime)
		util.HTTPInternalError(response, request, err)
		return
	}

	// TODO have WriteMultipartFile accept a multipart.File
	err = storage.WriteMultipartFile(id, file)
	if err != nil {
		log.Errorf("Could not write %s (mime %s) to disk", id, mime)
		util.HTTPInternalError(response, request, err)
		return
	}

	util.HTTPResponseJson(response, map[string]interface{}{"id": id}, 200)
	return
}

func RouteRoot(response http.ResponseWriter, request *http.Request) {
	log.Tracef("Routing root methods")

	switch request.Method {
	case "POST":
		handleUpload(response, request)
		return
	default:
		util.HTTPResponseError(response, "bad_method", 405)
		return
	}
}
