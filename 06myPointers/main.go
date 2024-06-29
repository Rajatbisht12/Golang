package main

import "fmt"

func main() {
	fmt.Println("Pointers in golang")

	// var ptr *int
	// fmt.Println(ptr)

	myNumber := 23

	var ptr = &myNumber //& refrencing

	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr * 2
	fmt.Println("Value of new Pointer is ", myNumber)
}
