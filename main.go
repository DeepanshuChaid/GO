package main

import (
	"fmt"
	"os"
)

func main() {
  file, err := os.Open("example.txt")
  if err != nil {
    fmt.Println(err)
  }

  fileInfo, err := file.Stat()
  if err != nil {
    panic(err)
  }
  fmt.Println(fileInfo.Size())


  os.Remove("example.txt")
  os.Remove("example1.2.txt")
  os.Remove("example_copy.txt")
}



