package server

import (
	"github.com/gastrodon/ferrothorn/storage"

	"net/http"
)

func respondErr(writer http.ResponseWriter, message string, code int) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(code)
	writer.Write([]byte(`{"error": "` + message + `"}`))
}

func ServeContent(writer http.ResponseWriter, request *http.Request, id string) {
	var path string
	var exists bool
	var err error
	if path, exists, err = storage.ReadPath(id); err != nil {
		respondErr(writer, "internal_error", 500)
		return
	}

	if !exists {
		respondErr(writer, "not_found", 404)
		return
	}

	if exists, err = storage.PathExists(path); err != nil {
		respondErr(writer, "internal_error", 500)
	}

	if !exists {
		go storage.DeleteID(id)
		respondErr(writer, "not_found", 404)
		return
	}

	http.ServeFile(writer, request, path)
	return
}
