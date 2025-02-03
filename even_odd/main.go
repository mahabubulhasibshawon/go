package main

import "fmt"

func myFunction(x int, y int) int {

	if x%2 == 0 {
		return x
	} else {
		return y
	}
}

func main() {
	fmt.Printf("even number is ")
	fmt.Println(myFunction(101, 2))
}
