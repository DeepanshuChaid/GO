package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/DeepanshuChaid/GO/internal/types"
)
  
func New() http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request) {

    var student types.Student;

    err := json.NewDecoder(r.Body).Decode(&student)
    if errors.Is(err, io.EOF) {
      slog.Error("Error decoding student", slog.String("error", err.Error()))
      w.WriteHeader(http.StatusBadRequest)
    }
    
    slog.Info("Creating a student")
    
    w.Write([]byte("Hello World"))
  }
}