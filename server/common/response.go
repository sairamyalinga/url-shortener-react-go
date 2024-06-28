package utils

import (
	"encoding/json"
	"net/http"
)

type JSONResponseData struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendJSONResponse(w http.ResponseWriter, data interface{}, code int, message string) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	response := JSONResponseData{
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(response)
}
