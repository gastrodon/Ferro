package routes

import (
	"monke-cdn/log"
	"monke-cdn/storage"
	"monke-cdn/util"

	"bytes"
	"crypto/md5"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
)

func getFileName(request *http.Request) (name string, err error) {
	name = request.URL.Query().Get("name")

	if name == "" {
		name, err = storage.NewUUID()
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

func Root(response http.ResponseWriter, request *http.Request) {
	defer request.Body.Close()
	log.Tracef("Handling upload")

	if !util.Authed(request.Header.Get("Authorization")) {
		log.Tracef("Auth header is bad")
		util.HTTPResponseError(response, "bad_auth", 401)
		return
	}

	if request.Method != "POST" {
		// TODO: route methods
		util.HTTPResponseError(response, "bad_method", 405)
		return
	}

	var file multipart.File
	var mime string
	var err error
	file, mime, err = getFile(request, "file", 32<<20)
	if err != nil {
		log.Tracef("getFile failed")
		util.HTTPResponseError(response, "bad_request", 400)
		return
	}

	var hash_sum []byte
	var hashable bytes.Buffer
	var file_dupe io.Reader = io.TeeReader(file, &hashable)
	hash_sum, err = md5File(hashable.Bytes())
	if err != nil {
		// TODO: don't hash during upload
		log.Tracef("hashing file failed")
		util.HTTPResponseError(response, "bad_request", 400)
		return
	}

	var id string
	id, err = getFileName(request)
	if err != nil {
		log.Tracef("getting file name failed")
		util.HTTPInternalError(response, request, err)
		return
	}

	err = storage.CreateReference(id, mime, hash_sum)
	if err != nil {
		log.Errorf("Creating a reference for %s (mime %s) failed", id, mime)
		util.HTTPInternalError(response, request, err)
		return
	}

	err = storage.WriteMultipartFile(id, file_dupe)
	if err != nil {
		log.Errorf("Writing %s (mime %s) to the disk failed", id, mime)
		util.HTTPInternalError(response, request, err)
		return
	}

	var r_map map[string]interface{} = map[string]interface{}{
		"id": id,
	}

	log.Tracef("File %s was uploaded", id)
	util.HTTPResponseJson(response, r_map, 200)
	return
}
