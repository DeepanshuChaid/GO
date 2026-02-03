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

  
  server.Shutdown()
  
  slog.Info("Server stopped")
  
}