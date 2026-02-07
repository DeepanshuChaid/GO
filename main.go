package main

import (
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
  "time"
  "context"
	"github.com/DeepanshuChaid/GO/internal/config"
      "github.com/DeepanshuChaid/GO/internal/http/handlers/student"
      "github.com/DeepanshuChaid/GO/internal/storage/sqlite"
)


func main () {

  cfg := config.MustLoad()

  storage, err := sqlite.New(cfg)
  if err != nil {
    log.Fatal(err)
  } 

  slog.Info("Storage initialized", slog.String("env", cfg.Env), slog.String("version", "1.0.0"))
  
  router := http.NewServeMux()

  router.HandleFunc("/api/students", student.New(storage))

  router.HandleFunc("/api/students/{id}", student.GetById(storage))

  server := http.Server{
    Addr: cfg.Address,
    Handler: router,
  }

  slog.Info("Starting server", slog.String("address", cfg.Address))
  fmt.Println("Server started on", cfg.Address)
  

  var done = make(chan os.Signal, 1)

  signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
  
  go func (){
    err := server.ListenAndServe()
    if err != nil {
      log.Fatal(err)
    }
  }()

  <-done


  ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
  defer cancel()
  fmt.Println(ctx)

  
  if err := server.Shutdown(ctx); err != nil {
    slog.Error("Server shutdown failed", slog.String("error", err.Error()))
  }
  
  slog.Info("Server stopped")
  
}