package storage

import (
	"monke-cdn/log"

	"fmt"
	"io"
	"os"
	"strings"
)

var file_root string

func SetFileRoot(where string) (err error) {
	log.Printf("Setting file root as %s", where)
	file_root = where
	if !strings.HasSuffix(file_root, "/") {
		file_root = fmt.Sprintf("%s/", file_root)
	}

	err = os.MkdirAll(where, os.ModePerm)
	return
}

func WriteMultipartFile(id string, file io.Reader) (err error) {
	var out *os.File
	var where string = fmt.Sprintf("%s%s", file_root, id)

	log.Printf("Writing file %s", where)
	out, err = os.Create(where)
	if err != nil {
		out.Close()
		return
	}

	io.Copy(out, file)
	out.Close()

	var info os.FileInfo
	info, err = os.Stat(where)
	if err != nil {
		log.Printf("Failed to get stat for %s", where)
		return
	}

	log.Printf("Wrote %d bytes", info.Size())
	return
}

func DeleteFile(path string) (err error) {
	log.Printf("Deleting path %s", path)
	err = os.Remove(path)
	return
}
