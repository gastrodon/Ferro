package main

import (
	"monke-cdn/server"
	"monke-cdn/storage"

	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

var (
	mongo_uname = os.Getenv("MONGO_USER")
	mongo_pass  = os.Getenv("MONGO_PASS")
	mongo_host  = os.Getenv("MONGO_HOST")
	db_name     = "monke-cdn"
)

func main() {
	var file_root *string = flag.String("at", "/monke/files/", "File storage root")
	var port *int = flag.Int("port", 8000, "port to serve")
	flag.Parse()

	var err error = storage.ConnectTo(fmt.Sprintf("%s:%s", mongo_uname, mongo_pass), mongo_host, db_name)
	err = storage.SetFileRoot(*file_root)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", server.RouteMain)

	log.Println("CDN online")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
