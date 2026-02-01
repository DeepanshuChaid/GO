package main

import (
	"fmt"
	"os"
  "bufio"
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
  
  // f.WriteString(" i like ahhhhhhhh wtf my does sukhmeen comes in my fkin mind")

  // bytes := []byte("i definitely like sukhmeen damn bitch ass ai really said that")
  // f.Write(bytes)

  fmt.Println("File created successfully")


  sourceFile, err := os.Open("example.txt")
  if err != nil {
     panic(err)
  }

  defer sourceFile.Close()

  destinationFile, err := os.Create("example_copy.txt")
  if err != nil {
     panic(err)
  }

  defer destinationFile.Close()

  reader := bufio.NewReader(sourceFile)
  writer := bufio.NewWriter(destinationFile)

  for {
    b, err := reader.ReadByte()
    if err != nil {
      if err.Error() == "EOF" {
        panic(err)
      }

      break
    }
    
    e := writer.WriteByte(b)
    if e != nil {
      panic(e)
    }
    
  }

  writer.Flush()


  fmt.Println("File copied successfully")
  
}


