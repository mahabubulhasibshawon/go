package main

import "fmt"

func myFunction(x int, y int) int {

	if x > y {
		return x - y
	} else {
		return x + y
	}
}

func main() {
	fmt.Println(myFunction(10, 2))
}
