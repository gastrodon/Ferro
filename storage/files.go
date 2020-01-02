package storage

import (
	"io"
	"os"
)

var file_root string

func SetFileRoot(where string) (err error) {
	file_root = where
	err = os.MkdirAll(where, os.ModePerm)
	return
}

func WriteMultipartFile(id string, file io.Reader) (err error) {
	var out *os.File
	defer out.Close()

	out, err = os.Create(file_root + id)
	if err != nil {
		return
	}

	io.Copy(out, file)
	return
}
