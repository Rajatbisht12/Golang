package main

import "fmt"

const Loggin string = "adfsjjjk"

func main() {
	var username string = "Rajat"
	fmt.Println(username)
	fmt.Printf("Variable is os type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is os type: %T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is os type: %T \n", smallVal)

	var smallFloat float32 = 255.564567467676
	fmt.Println(smallFloat)
	fmt.Printf("Variable is os type: %T \n", smallFloat)

	// implisit type

	var website = "hello"
	fmt.Println(website)

	// no var style

	noOfUser := 30000
	fmt.Println(noOfUser)

	fmt.Println(Loggin)
	fmt.Printf("Varibale type: %T \n", Loggin)
}
