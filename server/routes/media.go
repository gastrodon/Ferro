package routes

import (
	"monke-cdn/log"
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

func deleteMedia(response http.ResponseWriter, request *http.Request, result map[string]interface{}) {
	log.Tracef("Deleting file %s (at %s)", result["id"].(string), result["path"].(string))

	var err error
	_, err = storage.DeleteUnique(bson.D{{"id", result["id"].(string)}})
	if err != nil {
		log.Errorf("Deleting reference to %s failed", result["id"].(string))
		util.HTTPInternalError(response, request, err)
		return
	}

	err = storage.DeleteFile(result["path"].(string))
	if err != nil {
		log.Errorf("Deleting file %s failed", result["path"].(string))
		util.HTTPInternalError(response, request, err)
		return
	}

	log.Tracef("Deleted file %s (at %s)", result["id"].(string), result["path"].(string))
	response.WriteHeader(204)
	return
}

func Media(response http.ResponseWriter, request *http.Request) {
	var id string = strings.Split(request.URL.Path, ".")[0][1:]
	log.Tracef("Handling media path for id %s", id)

	var result map[string]interface{}
	var exists bool
	var err error
	result, exists, err = storage.GetUnique(bson.D{{"id", id}})
	if err != nil {
		log.Errorf("Could not get a unique of %s", id)
		util.HTTPInternalError(response, request, err)
		return
	}

	if !exists {
		log.Tracef("A reference to %s was not found", id)
		util.HTTPResponseError(response, "not_found", 404)
		return
	}

	switch request.Method {
	case "GET":
		uploadMedia(response, request, result)
		return
	case "DELETE":
		deleteMedia(response, request, result)
		return
	}

	log.Tracef("Bad method %s", request.Method)
	util.HTTPResponseError(response, "bad_method", 405)
}
