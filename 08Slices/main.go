package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitlist = []string{"Apple", "Peache", "Tomato"}

	fmt.Printf("%T\n", fruitlist)
	fmt.Println(fruitlist)

	fruitlist = append(fruitlist, "Mango")

	fmt.Println(fruitlist)

	fruitlist = append(fruitlist[1:])
	fmt.Println(fruitlist)

	highScore := make([]int, 4)
	highScore[0] = 235
	highScore[1] = 945
	highScore[2] = 455
	highScore[3] = 500

	fmt.Println(highScore)

	highScore = append(highScore, 555, 234, 667)
	fmt.Println(highScore)

	sort.Ints(highScore)
	fmt.Println(highScore)
	fmt.Println(sort.IntsAreSorted(highScore))
}
