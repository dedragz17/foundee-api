package controller

import (
	"encoding/json"
	"net/http"
)

func IndexController(writer http.ResponseWriter, request *http.Request) {

	response := make(map[string]interface{})
	response["version"] = "1.0"

	json.NewEncoder(writer).Encode(response)
	writer.WriteHeader(http.StatusOK)
}
