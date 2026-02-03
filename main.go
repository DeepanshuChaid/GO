package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"

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

  var done = make(chan os.Signal, 1)

  signal.Notify(done, os.Interrupt, sysca)
  
  go func (){
    err := server.ListenAndServe()
    if err != nil {
      log.Fatal(err)
    }
  }()

  <-done
}