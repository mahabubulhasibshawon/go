package main

import (
	"fmt"
	"time"
)

func main() {
	switch time.Now().Weekday() {
	case time.Friday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's weekday")
	}

	t := time.Now()
	fmt.Println(t.Local().Date())
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Print("It's afternoon")
	}
}
