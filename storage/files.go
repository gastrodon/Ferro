package storage

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var file_root string

func SetFileRoot(where string) (err error) {
	file_root = where
	if !strings.HasSuffix(file_root, "/") {
		file_root = fmt.Sprintf("%s/", file_root)
	}

	err = os.MkdirAll(where, os.ModePerm)
	return
}

func WriteMultipartFile(id string, file io.Reader) (err error) {
	var out *os.File
	defer out.Close()

	out, err = os.Create(fmt.Sprintf("%s%s", file_root, id))
	if err != nil {
		return
	}

	io.Copy(out, file)
	return
}
