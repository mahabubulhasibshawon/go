package main

import "fmt"

func print(numbers ...int) {
	fmt.Println(numbers)
	fmt.Println("Length : ", len(numbers))
	fmt.Println("Capacity : ", cap(numbers))
}

func main() {
	print(1, 2, 3, 4, 5)
}
