package server

import (
	"fmt"
	"monke-cdn/server/routes"
	"monke-cdn/util"
	"net/http"
	"regexp"
)

var filename_pattern string = "(.)+((.).+)?"
var root_pattern = regexp.MustCompile("^/(&.+)?$")
var content_pattern = regexp.MustCompile(fmt.Sprintf("^/%s/?(&.+)?$", filename_pattern))
var md5_pattern = regexp.MustCompile(fmt.Sprintf("^/%s/md5/?(&.+)$", filename_pattern))
var thumb_pattern = regexp.MustCompile(fmt.Sprintf("^/%s/thumb/?(&.+)$", filename_pattern))

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
