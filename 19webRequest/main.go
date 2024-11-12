package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://www.google.com"

func main(){
  fmt.Println("Hello")

  res, err := http.Get(url)

  if err != nil{
    panic(err)
  }

  fmt.Printf("Response is of type: %T\n", res)
  defer res.Body.Close()

  databytes, err := ioutil.ReadAll(res.Body)

  if err != nil{
    panic(err)
  }

  content := string(databytes)

  fmt.Println(content)
}