package server

import (
	"fmt"
	"log"
	"monke-cdn/server/routes"
	"monke-cdn/util"
	"net/http"
	"regexp"
)

var root_pattern *regexp.Regexp
var content_pattern *regexp.Regexp
var md5_pattern *regexp.Regexp
var thumb_pattern *regexp.Regexp

func BuildRoutes() {
	var regexp_err error
	var filename_pattern string = "(.)+((.).+)?"

	root_pattern, regexp_err = regexp.Compile("^/(&.+)?$")
	content_pattern, regexp_err = regexp.Compile(fmt.Sprintf("^/%s/?(&.+)?$", filename_pattern))
	md5_pattern, regexp_err = regexp.Compile(fmt.Sprintf("^/%s/md5/?(&.+)$", filename_pattern))
	thumb_pattern, regexp_err = regexp.Compile(fmt.Sprintf("^/%s/thumb/?(&.+)$", filename_pattern))

	if regexp_err != nil {
		log.Fatal(regexp_err)
	}
}

func RouteMain(response http.ResponseWriter, request *http.Request) {
	var r_map map[string]interface{} = map[string]interface{}{
		"path": request.URL.Path,
	}

	var path string = request.URL.Path

	switch {
	case root_pattern.MatchString(path):
		routes.Root(response, request)
		return
	case content_pattern.MatchString(path):
		routes.Media(response, request)
		return
	case md5_pattern.MatchString(path):
		fmt.Println("md5_pattern")
		return
	case thumb_pattern.MatchString(path):
		fmt.Println("thumb_pattern")
		return
	}

	util.HTTPResponseJson(response, r_map, 200)
}
