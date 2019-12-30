package main

import (
	"fmt"
	"log"
	"monke-cdn/server"
	"net/http"
)

func main() {
	server.RecieveSecret("foobar")
	server.BuildRoutes()
	http.Handle("/", server.BuildHandler(
		http.HandlerFunc(server.RouteMain),
		server.AuthenticateRoute,
	))
	fmt.Println("CDN online")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":8000"), nil))
}
