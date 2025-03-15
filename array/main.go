package main

import "fmt"

func main() {
	myArray := [...]int{1, 2, 3}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(i)
	}
}
