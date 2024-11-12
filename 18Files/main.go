package main

import (
	"fmt"
	"io"
	"io/ioutil"
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
  defer file.Close()
  readFile("./myfile.txt")
}

func readFile(filename string){
  databyte, err := ioutil.ReadFile(filename)
  if err != nil{
    panic(err)
  }
  fmt.Println(string(databyte))
}