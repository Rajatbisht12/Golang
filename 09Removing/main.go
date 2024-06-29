package main

import "fmt"

func main() {
	var courses = []string{"React", "CPP", "Swift", "Java"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
