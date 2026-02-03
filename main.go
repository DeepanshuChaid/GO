package main

import (
	"fmt"
	"net/http"
    "log"
	"github.com/DeepanshuChaid/GO/internal/config"
)


func main () {

  cfg := config.MustLoad()


  router := http.NewServeMux()

  router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World")
  })


  server := http.Server{
    Addr: cfg.Address,
    Handler: router,
  }

  err := server.ListenAndServe()
  if err != nil {
    log.Fatal(err)
  }
}