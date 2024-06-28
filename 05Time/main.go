package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time in Golang")

	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))

	createDate := time.Date(2024, time.August, 10, 12, 56, 55, 0, time.UTC)
	fmt.Println(createDate)
}
