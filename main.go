package main

import (
	"monke-cdn/log"
	"monke-cdn/server"
	"monke-cdn/storage"

	"flag"
	"fmt"
	"net/http"
	"os"
)

var (
	mongo_connection = os.Getenv("FERRO_CONNECTION")
	db_name          = os.Getenv("FERRO_MONGO_BASE")
)

func main() {
	var (
		level     *int    = flag.Int("level", 1, "logging level")
		file_root *string = flag.String("at", "/monke/files/", "File storage root")
		port      *int    = flag.Int("port", 8000, "port to serve")
	)
	flag.Parse()

	log.At(*level)

	var err error = storage.ConnectTo(mongo_connection, db_name)
	err = storage.SetFileRoot(*file_root)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/", server.RouteMain)

	log.Println("CDN online")
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *port), nil))
}
