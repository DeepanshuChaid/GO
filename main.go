package main

import (
  "fmt"
  "time"
)

func fetchData() string {
  return "data"
}

func main() {
    resultChan := make(chan int)

    go func() {
        fmt.Println("Starting slow task...")
        time.Sleep(2 * time.Second)
        resultChan <- 42
        fmt.Println("Sent result!")
    }()

    fmt.Println("Before waiting...")

    result := <-resultChan  // ⛔ BLOCKS HERE! WAITS!

    fmt.Println("Got result:", result)  // Won't run until channel has data
}
