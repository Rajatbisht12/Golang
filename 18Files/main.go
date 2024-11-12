package main

import (
	"fmt"
	"io"
	"os"
)

func main(){
  fmt.Println("Hello")
  content := "This needs to go in a file - content"
  file, err := os.Create("./myfile.txt")

  if err != nil{
    panic(err)
  }

  length, err := io.WriteString(file, content)
  if err != nil{
    panic(err)
  }
  fmt.Println(length)
  file.Close()
}