package main 

import (
  "fmt"
  "github.com/DeepanshuChaid/GO/internal/config"
)


func main () {

  cfg := config.MustLoad()
  
  
  fmt.Println("Hello World")
}