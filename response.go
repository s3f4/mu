package mu

import (
	"encoding/json"
	"net/http"
)

// Response is used to create an http response
type Response struct {
	Status  bool                   `json:"status"`
	Message string                 `json:"message,omitempty"`
	Data    map[string]interface{} `json:"data,omitempty"`
}

//SendResponse returns json response
func SendResponse(w http.ResponseWriter, statusCode int, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	var response *Response = dataOrMessage(message)
	response.Status = statusCode == http.StatusOK
	resp, err := json.Marshal(response)

	if err != nil {
		return
	}

	w.Write(resp)
}

func dataOrMessage(message interface{}) *Response {
	r := &Response{}
	switch message.(type) {
	case string:
		r.Message = message.(string)
	case map[string]interface{}:
		r.Data = message.(map[string]interface{})
	}
	return r
}

//R200 returns Status ok
func R200(w http.ResponseWriter, message interface{}) {
	SendResponse(w, http.StatusOK, message)
}

//R400 returns Status bad request
func R400(w http.ResponseWriter, message interface{}) {
	SendResponse(w, http.StatusBadRequest, message)
}

//R401 returns Status unauthorized
func R401(w http.ResponseWriter, message interface{}) {
	SendResponse(w, http.StatusUnauthorized, message)
}

//R403 returns Status forbidden
func R403(w http.ResponseWriter, message interface{}) {
	SendResponse(w, http.StatusForbidden, message)
}

//R404 returns Status not found
func R404(w http.ResponseWriter, message interface{}) {
	SendResponse(w, http.StatusNotFound, message)
}

//R422 returns Status unprocessable entity
func R422(w http.ResponseWriter, message interface{}) {
	SendResponse(w, http.StatusUnprocessableEntity, message)
}

//R500 returns Status internal server error
func R500(w http.ResponseWriter, message interface{}) {
	SendResponse(w, http.StatusInternalServerError, message)
}
