package main

import "fmt"

func main() {
	greeter()

	result := adder(3, 5)

	fmt.Println("The value is ", result)

	result1, message := proAdder(2, 4, 5, 6, 7, 9, 35)
	fmt.Println(result1, message)
}

func adder(valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "Hello"
}

func greeter() {
	fmt.Println("There is something")
}
