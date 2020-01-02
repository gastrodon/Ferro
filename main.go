package main

import (
	"monke-cdn/server"
	"monke-cdn/storage"
	"monke-cdn/util"

	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	var file_root *string = flag.String("at", "/monke/files/", "File storage root")
	flag.Parse()

	var err error = storage.ConnectTo(fmt.Sprintf("%s:%s", mongo_uname, mongo_pass), mongo_host, "monke-cdn")
	err = storage.SetFileRoot(*file_root)
	if err != nil {
		log.Fatal(err)
	}

	util.RecieveSecret(secret)
	server.BuildRoutes()
	http.HandleFunc("/", server.RouteMain)

	fmt.Println("CDN online")

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":8000"), nil))
}
