package server

import (
	"github.com/gastrodon/ferrothorn/storage"
	"github.com/google/uuid"

	"mime/multipart"
	"net/http"
	"strings"
)

const (
	FILE_PART     = "file"
	MULTIPART_MEM = 32 << 20
)

func splitIgnoreEmpty(it rune) (ok bool) {
	ok = it == '/'
	return
}

func UploadContent(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	code, r_map, err = upload(request, uuid.New().String())
	return
}

func UploadNamedContent(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var name string = strings.FieldsFunc(request.URL.Path, splitIgnoreEmpty)[0]
	code, r_map, err = upload(request, name)
	return
}

func getFile(request *http.Request) (file multipart.File, mime string, err error) {
	if err = request.ParseMultipartForm(MULTIPART_MEM); err != nil {
		return
	}

	var header *multipart.FileHeader
	// TODO: test to see if file can just be populated here
	//       and if I can skip `header.Open()`
	if _, header, err = request.FormFile(FILE_PART); err != nil || header == nil {
		return
	}

	mime = header.Header.Get("Content-Type")
	file, err = header.Open()
	return
}

func upload(request *http.Request, name string) (code int, r_map map[string]interface{}, err error) {
	var file multipart.File
	var mime string
	if file, mime, err = getFile(request); err != nil {
		err = nil
		code = 400
		return
	}

	if err = storage.CreateFile(name, mime, file); err != nil {
		return
	}

	code = 200
	r_map = map[string]interface{}{"id": name}
	return
}

func DeleteContent(request *http.Request) (code int, r_map map[string]interface{}, err error) {
	var id string = strings.FieldsFunc(request.URL.Path, splitIgnoreEmpty)[0]

	var path string
	if path, _, err = storage.ReadPath(id); err != nil {
		return
	}

	go storage.DeleteID(id)
	go storage.DeleteFile(path)
	return
}
