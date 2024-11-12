package main

import "fmt"

func main(){
  defer fmt.Println("World")
  defer fmt.Println("One")
  defer fmt.Println("Two")
  fmt.Println("Hello")
  // The deffer works in LIFO and it will print after print function

  myDefer()
  fmt.Println()
}

func myDefer(){
  for i := 0; i < 5; i++{
    defer fmt.Print(i)
  }
}