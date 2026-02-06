package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"
  // "github.com/DeepanshuChaid/GO/internal/storage"

	"github.com/DeepanshuChaid/GO/internal/http/utils/response"
	"github.com/DeepanshuChaid/GO/internal/types"
	"github.com/go-playground/validator/v10"

  "github.com/DeepanshuChaid/GO/internal/storage"
)
  
func New(storage storage.Storage) http.HandlerFunc{
  return func(w http.ResponseWriter, r *http.Request) {

    var student types.Student;

    // err := json.NewDecoder(r.Body).Decode(&student)
    err := json.NewDecoder(r.Body).Decode(&student)
    if errors.Is(err, io.EOF) {
      response.WriteJson(w, http.StatusBadRequest, response.GenralError(fmt.Errorf("Request body is empty")))
      return
    }

    if err != nil {
      response.WriteJson(w, http.StatusBadRequest, response.GenralError(err))
    }
    
    // request validation 
    
    if err := validator.New().Struct(student); err != nil {
      validate := err.(validator.ValidationErrors)
      
      response.WriteJson(w, http.StatusBadRequest, response.ValidateError(validate))
      return
    }

    // id, err := storage.CreateStudent(
    //   student.Name,
    //   student.Email,
    //   student.Age,
    // )
    
    slog.Info("Creating a student")

    response.WriteJson(w, http.StatusOK, map[string]string{"message": "Student created"})
    
    // w.Write([]byte("Hello World"))
  }
}

