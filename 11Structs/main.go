package main

import "fmt"

func main() {
	fmt.Println("This is an struct video")

	myUser := User{"Rajat", "bisrajat123@gmail.com", true, 23}
	fmt.Println(myUser)
	fmt.Printf("%+v\n", myUser)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
