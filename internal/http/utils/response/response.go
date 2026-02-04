package response

import (
  "net/http"
)

func WriteJson(w http.ResponseWriter, status int, data interface {}) {
  w.Header().Set("Content-Type", "application/json")
  w.WriteHeader(status)
}