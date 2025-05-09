package utils

// DoubleSlice returns a new slice with each element doubled.
func DoubleSlice(input []int) []int {
	doubled := make([]int, len(input))
	for i, value := range input {
		doubled[i] = value * 2
	}
	return doubled
}

// FilterEven returns a new slice containing only the even numbers from the input slice.
func FilterEven(input []int) []int {
	evenNumbers := []int{}
	for _, value := range input {
		if value%2 == 0 {
			evenNumbers = append(evenNumbers, value)
		}
	}
	return evenNumbers
}

// ReverseSlice returns a reversed copy of the input slice
func ReverseSlice(input []int) []int {
	reversed := make([]int, len(input))
	for i, value := range input {
		reversed[len(input)-1-i] = value
	}
	return reversed
}

func RemoveSpecificElement(input []int, element int) {
	removedSlice := make([]int, len(input)-1)
	for _, value := range input {
		if value ==  element{
			continue
		}
		// removedSlice := append(removedSlice , input...)
	}
	// return removedSlice
}