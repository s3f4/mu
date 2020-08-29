package helpers

import (
	"encoding/json"
	"net/http"
)

// Response is used to create an http response
type Response struct {
	Status  bool                   `json:"status"`
	Message string                 `json:"message,omitempty"`
	Error   interface{}            `json:"error,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

//SendResponse returns json response
func SendResponse(w http.ResponseWriter, statusCode int, response Response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	if statusCode != http.StatusOK {
		response.Status = false
	} else {
		response.Status = true
	}

	if response.Error != nil {
		response.Error = response.Error.(error).Error()
	}

	resp, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(resp)
}

//R200 returns Status ok
func R200(w http.ResponseWriter, response Response) {
	SendResponse(w, http.StatusOK, response)
}

//R400 returns Status bad request
func R400(w http.ResponseWriter, response Response) {
	SendResponse(w, http.StatusBadRequest, response)
}

//R401 returns Status unauthorized
func R401(w http.ResponseWriter, response Response) {
	SendResponse(w, http.StatusUnauthorized, response)
}

//R403 returns Status not found
func R403(w http.ResponseWriter, response Response) {
	SendResponse(w, http.StatusForbidden, response)
}

//R404 returns Status not found
func R404(w http.ResponseWriter, response Response) {
	SendResponse(w, http.StatusNotFound, response)
}

//R422 returns Status unprocessable entity
func R422(w http.ResponseWriter, response Response) {
	SendResponse(w, http.StatusUnprocessableEntity, response)
}

//R500 returns Status internal server error
func R500(w http.ResponseWriter, response Response) {
	SendResponse(w, http.StatusInternalServerError, response)
}
