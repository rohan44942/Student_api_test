package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"status"`
	Error  string `json:"error"`
}

const (
	statusOk    = "OK"
	statusError = "ERROR"
)

func WriteJson(w http.ResponseWriter, status int, data any) error {
	w.Header().Set("content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)

}

func GeneralError(err error) Response {
	return Response{
		Status: statusError,
		Error:  err.Error(),
	}
}
