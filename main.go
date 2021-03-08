package main

import (
	"github.com/gastrodon/ferrothorn/server"
	"github.com/gastrodon/ferrothorn/storage"

	"github.com/gastrodon/groudon/v2"

	"log"
	"net/http"
	"os"
	"strings"
)

var (
	root = os.Getenv("FERROTHORN_ROOT")
)

const (
	FILENAME_PATTERN = `[a-zA-Z0-9\-\.]+`

	ROUTE_ANY  = `^/.*$`
	ROUTE_ROOT = `^/$`
	ROUTE_FILE = `^/` + FILENAME_PATTERN + `/?$`
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

	groudon.AddMiddleware("POST", ROUTE_ANY, server.MustAuth)
	groudon.AddMiddleware("DELETE", ROUTE_ANY, server.MustAuth)

	groudon.AddHandler("POST", ROUTE_ROOT, server.UploadContent)
	groudon.AddHandler("POST", ROUTE_FILE, server.UploadNamedContent)
	groudon.AddHandler("DELETE", ROUTE_FILE, server.DeleteContent)
	http.Handle("/", http.HandlerFunc(splitter))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
