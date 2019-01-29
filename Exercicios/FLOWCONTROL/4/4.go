package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Should I go to work today?")

	weekday := time.Now().Weekday().String()

	switch weekday {
	case "Saturday":
		fallthrough
	case "Sunday":
		fmt.Println("No.")
	default:
		fmt.Println("Yes.")
	}
}
