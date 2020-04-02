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
	log.Tracef("Setting file root as %s", where)
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

	log.Tracef("Writing file %s", where)
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
		log.Errorf("Failed to get stat for %s", where)
		return
	}

	log.Tracef("Wrote %d bytes", info.Size())
	return
}

func DeleteFile(path string) (err error) {
	log.Tracef("Deleting path %s", path)
	err = os.Remove(path)
	return
}

func Exists(path string) (exists bool, err error) {
	var info os.FileInfo
	info, err = os.Stat(path)

	if err == nil {
		exists = !info.IsDir()
		return
	}

	exists = false
	if os.IsNotExist(err) {
		err = nil
	}

	return
}
