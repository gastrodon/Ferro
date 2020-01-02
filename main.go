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
	var csr *string = flag.String("csr", "", "csr to use for SSL")
	var key *string = flag.String("key", "", "key to use for SSL")
	var port *int = flag.Int("p", 8000, "port to serve")
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

	if len(*csr)+len(*key) != 0 {
		log.Fatal(http.ListenAndServeTLS(fmt.Sprintf(":%d", *port), *csr, *key, nil))
	}
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
