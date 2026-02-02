package auth

import (
  "fmt"
)

type User struct {
  Name string
  Email string
  Password string
}

func Login () {
  fmt.Println("login")
}