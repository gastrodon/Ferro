package util

import (
	"monke-cdn/log"

	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func HTTPResponseJson(response http.ResponseWriter, response_map map[string]interface{}, code int) {
	var response_data []byte
	var parse_err error
	response_data, parse_err = json.Marshal(response_map)

	if parse_err != nil {
		response_data = []byte(`{"error": "internal_error"}`)
		code = 500
	}

	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(code)
	fmt.Fprintf(response, string(response_data))
}

func HTTPResponseError(response http.ResponseWriter, err_message string, code int) {
	var response_map map[string]interface{} = map[string]interface{}{
		"error": err_message,
	}

	HTTPResponseJson(response, response_map, code)
}

func HTTPInternalError(response http.ResponseWriter, request *http.Request, err error) {
	var stamp int64 = time.Now().Unix()

	log.Errorf("[%d] %s %s failed! \n[body] %s\n[error] %s", stamp/1000, request.Method, request.URL.Path, request.Body, err.Error())

	HTTPResponseError(response, "internal_error", 500)
}
