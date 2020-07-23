package server

import (
	"net/http"
	"os"
)

var (
	secret = os.Getenv("FERROTHORN_SECRET")
)

func MustAuth(request *http.Request) (_ *http.Request, ok bool, code int, r_map map[string]interface{}, err error) {
	var header string = request.Header.Get("Authorization")
	ok = header == "" || header == secret
	code = 401
	return
}
