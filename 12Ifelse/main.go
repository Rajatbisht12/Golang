package main

import "fmt"

func main() {
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "less than"
	} else if loginCount > 10 {
		result = "more than"
	} else {
		result = "exactly"
	}

	fmt.Println(result)
}
