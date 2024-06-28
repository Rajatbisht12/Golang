package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")

	// comma ok || err ok

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating ", input)
}
