package main

import (
        "fmt"
)

func main() {
        const name = "nigga"
        const age = 25

        msg := fmt.Sprintf("Hello, %s! You are %d years old.", name, age)
        
        fmt.Println(msg)
}
