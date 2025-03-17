package main

import "fmt"

func sliceExample() {
	s := []byte{1, 2}

	fmt.Println(s)
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	sliceExample()
	fmt.Println(numbers)
	fmt.Printf("numbers has a data type of : %T\n", numbers)

}
