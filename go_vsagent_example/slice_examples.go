package main

import "fmt"

func main() {
	// Creating a slice
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println("Original slice:", numbers)

	// Appending to a slice
	numbers = append(numbers, 6, 7)
	fmt.Println("After appending:", numbers)

	// Slicing a slice
	subSlice := numbers[1:4]
	fmt.Println("Sub-slice (1:4):", subSlice)

	// Modifying a slice
	numbers[0] = 10
	fmt.Println("After modifying the first element:", numbers)

	// Length and capacity of a slice
	fmt.Println("Length:", len(numbers))
	fmt.Println("Capacity:", cap(numbers))

	// Creating a slice with make
	newSlice := make([]int, 3, 5)
	fmt.Println("Slice created with make:", newSlice)
	fmt.Println("Length:", len(newSlice))
	fmt.Println("Capacity:", cap(newSlice))

	// Copying a slice
	copySlice := make([]int, len(numbers))
	copy(copySlice, numbers)
	fmt.Println("Copied slice:", copySlice)
}
