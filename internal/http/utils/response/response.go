package response

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Status string `json:"custom_status"`
	Error  string 
}

const (
  StatusOK = "OK"
  StatusError = "Error"
)

func WriteJson(w http.ResponseWriter, status int, data interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(data)
}



func GenralError(err error) {
  return Response{
    Status: StatusError,
    Error: err.Error(),
  }
}