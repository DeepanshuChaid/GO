package main

import (
	"fmt"
	"github.com/fatih/color"

	"github.com/DeepanshuChaid/GO/auth"
)

func main (){
  fmt.Println("hello world")

  auth.Login()

  user1 := auth.User{
    Name: "Deepanshu",
    Email: "deepanshu@chaid.com",
    Password: "12345",
  }

  fmt.Println(user1)

  color.Red(user1.Email)

  color.RGB(255, 128, 0).Println("foreground orange")
  color.RGB(230, 42, 42).Println("foreground red")

  color.BgRGB(255, 128, 0).Println("background orange")
  color.BgRGB(230, 42, 42).Println("background red")
}