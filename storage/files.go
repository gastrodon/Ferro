package storage

import (
	"io"
	"mime/multipart"
	"os"
	"strings"
)

var (
	file_root string
)

func FileRoot(path string) (err error) {
	file_root = strings.TrimSuffix(path, "/") + "/"
	err = os.MkdirAll(file_root, os.ModePerm)
	return
}

func WriteMultipartFile(id string, file multipart.File) (path string, err error) {
	path = file_root + id

	var out *os.File
	if out, err = os.Create(path); err != nil {
		return
	}

	io.Copy(out, file)
	out.Close()
	return
}

func DeleteFile(path string) (err error) {
	err = os.Remove(path)
	return
}

func PathExists(path string) (exists bool, err error) {
	var info os.FileInfo
	if info, err = os.Stat(path); err == nil {
		exists = !info.IsDir()
		return
	}

	if os.IsNotExist(err) {
		err = nil
	}

	return
}
