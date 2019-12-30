package server

import (
	"monke-cdn/util"
	"net/http"
)

var secret string

func RecieveSecret(recieved string) {
	secret = recieved
}

func BuildHandler(handler http.Handler, functions ...func(http.Handler) http.Handler) http.Handler {
	var stacked func(http.Handler) http.Handler

	for _, stacked = range functions {
		handler = stacked(handler)
	}

	return handler
}

func AuthenticateRoute(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(response http.ResponseWriter, request *http.Request) {
		var auth string = request.Header.Get("Authorization")

		if auth == secret {
			util.HTTPResponseError(response, "bad_auth", 401)
			return
		}

		handler.ServeHTTP(response, request)
	})
}
