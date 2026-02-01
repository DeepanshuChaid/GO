package main

import (
	"fmt"
	"os"
)

func main() {
  // f, err := os.ReadFile("example.txt")
  // if err != nil {
  //   panic(err)
  // }

  // fmt.Println(string(f))

  // dir, err := os.Open(".")
  // if err != nil {
  //   panic(err)
  // }

  // defer dir.Close()

  // fileInfo, err := dir.Readdir(-1)

  // for _, file := range fileInfo {
  //   fmt.Println(file.Name())
  // }


  // array := []int{1,2,3,4,5,6,7,8,9,10}

  // for i, v := range array {
  //   fmt.Println(i, v)
  // }

  f, err := os.Create("example1.2.txt")
  if err != nil {
    panic(err)
  }
  
  defer f.Close()
  
  f.WriteString(" i like ahhhhhhhh wtf my does sukhmeen comes in my fkin mind")

fmt.Println("File created successfully")
}
