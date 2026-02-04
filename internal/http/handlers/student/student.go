package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/DeepanshuChaid/GO/internal/http/utils/response"
	"github.com/DeepanshuChaid/GO/internal/types"
)
  
func New() http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request) {

    var student types.Student;

    err := json.NewDecoder(r.Body).Decode(&student)
    if errors.Is(err, io.EOF) {
      response.WriteJson(w, http.StatusBadRequest, map[string]string{"error": "request body is empty"})
      return
    }
    
    slog.Info("Creating a student")

    response.WriteJson(w, http.StatusOK, map[string]string{"message": "Student created"})
    
    w.Write([]byte("Hello World"))
  }
}