package main

import "fmt"

func main() {
	days := []string{"Sunday", "Monday", "Saturday", "Tuesday", "Friday"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	rougheValue := 1
	for rougheValue < 10 {

		if rougheValue == 5 {
			rougheValue++
			continue
		}

		fmt.Println("Value is ", rougheValue)
		rougheValue++
	}
}
