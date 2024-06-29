package main

import (
	"fmt"
)

func main() {
	fmt.Println("Maps in go lang")
	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println(languages)
	delete(languages, "JS")

	// For loop in GOLang

	for key, value := range languages {
		fmt.Printf("key is %v and the value is %v\n", key, value)
	}
}
