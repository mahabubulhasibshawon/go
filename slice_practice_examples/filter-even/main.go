package main

import (
	"fmt"
	"slice_practice_examples/filter-even/utils"
	// Import the utils package
)

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	doubledNumbers := utils.DoubleSlice(numbers) // Call the function from the utils package
	fmt.Println("Original Numbers:", numbers)
	fmt.Println("Doubled Numbers:", doubledNumbers)

	evenNumbers := utils.FilterEven(numbers)
	fmt.Println("Even Numbers:", evenNumbers)

	reversedNumbers := utils.ReverseSlice(numbers)
	fmt.Println("Reversed Numbers:", reversedNumbers)

	// Important Note on Slices:
	// Slices in Go are references to underlying arrays.  When you pass a slice
	// to a function, you're passing a reference.  If the function modifies the
	// *elements* of the slice (e.g., by doing `input[0] = 100`), those changes
	// *will* be seen by the caller.  However, if the function changes the slice
	// *itself* (e.g., by assigning a new slice to the `input` variable, or
	// appending in a way that causes a new allocation), those changes will *not*
	// be seen by the caller.  The examples above avoid modifying the original
	// slice, and instead create and return *new* slices.
}