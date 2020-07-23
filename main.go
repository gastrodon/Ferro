package main

import (
	"github.com/gastrodon/ferrothorn/server"
	"github.com/gastrodon/ferrothorn/storage"

	"github.com/gastrodon/groudon"

	"log"
	"net/http"
	"os"
	"strings"
)

var (
	root = os.Getenv("FERROTHORN_ROOT")
)

func splitIgnoreEmpty(it rune) (ok bool) {
	ok = it == '/'
	return
}

func splitter(writer http.ResponseWriter, request *http.Request) {
	var split []string = strings.FieldsFunc(request.URL.Path, splitIgnoreEmpty)

	switch {
	case request.Method != "GET", len(split) != 1:
		groudon.Route(writer, request)
	default:
		server.ServeContent(writer, request, split[0])
	}

	return
}

func main() {
	storage.Connect(os.Getenv("FERROTHORN_CONNECTION"))
	storage.FileRoot(root)

	groudon.RegisterMiddleware(server.MustAuth)
	groudon.RegisterHandler("POST", `^/$`, server.UploadContent)
	groudon.RegisterHandler("POST", `^/[a-zA-Z0-9\-\.]+/?$`, server.UploadNamedContent)
	groudon.RegisterHandler("DELETE", `^/[a-zA-Z0-9\-\.]+/?$`, server.DeleteContent)
	http.Handle("/", http.HandlerFunc(splitter))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
