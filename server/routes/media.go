package routes

import (
	"monke-cdn/storage"
	"monke-cdn/util"

	"go.mongodb.org/mongo-driver/bson"
	"net/http"
	"strings"
)

func uploadMedia(response http.ResponseWriter, request *http.Request, result map[string]interface{}) {
	http.ServeFile(response, request, result["path"].(string))
	return
}

func Media(response http.ResponseWriter, request *http.Request) {
	var id string = strings.Split(request.URL.Path, ".")[0][1:]

	var result map[string]interface{}
	var exists bool
	var err error
	result, exists, err = storage.GetUnique(bson.D{{"id", id}})
	if err != nil {
		util.HTTPInternalError(response, request, err)
		return
	}

	if !exists {
		util.HTTPResponseError(response, "not_found", 404)
		return
	}

	switch request.Method {
	case "GET":
		uploadMedia(response, request, result)
		return
	}

	util.HTTPResponseError(response, "bad_method", 405)
}
