package main

import (
	"fmt"
  "net/http"
  "log"
	"github.com/DeepanshuChaid/Go/internal/config"
)

func main() {

  cfg := config.MustLoad()


  router := http.NewServeMux()

  router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("welcome to the home page"))
  })

  server := http.Server {
    Addr: cfg.HttpServer.Address,
    Handler: router,
  }

  fmt.Println("Server is running on", cfg.HttpServer.Address)

  err := server.ListenAndServe()
  if err != nil {
    log.Fatalf("server failed to start: %v", err)
  }
  
}