package main

import (
	"monke-cdn/server"
	"monke-cdn/storage"

	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var file_root *string = flag.String("at", "/monke/files/", "File storage root")
	flag.Parse()

	var err error = storage.ConnectTo(fmt.Sprintf("%s:%s", mongo_uname, mongo_pass), mongo_host, "monke-cdn")
	fmt.Println(*file_root)
	err = storage.SetFileRoot(*file_root)
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
