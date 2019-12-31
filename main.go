package main

import (
	"fmt"
	"log"
	"monke-cdn/server"
	"monke-cdn/storage"
	"net/http"
)

func main() {
	var err error
	err = storage.ConnectTo(fmt.Sprintf("%s:%s", mongo_uname, mongo_pass), mongo_host, "monke-cdn")

	if err != nil {
		log.Fatal(err)
	}

	server.RecieveSecret(secret)
	server.BuildRoutes()
	http.Handle("/", server.BuildHandler(
		http.HandlerFunc(server.RouteMain),
		server.AuthenticateRoute,
	))
	fmt.Println("CDN online")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":8000"), nil))
}
