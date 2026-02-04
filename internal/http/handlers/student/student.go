package student

import (
  "net/http"
  "log/slog"
)
  
func New() http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request) {
    
    slog.Info("Creating a student")
    
    w.Write([]byte("Hello World"))
  }
}