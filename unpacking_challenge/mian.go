package main

import "fmt"

func addAll(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum
}

func avgAll(nums ...int) int {
	sum := 0
	for _, n := range nums {
		sum += n
	}
	return sum / len(nums)
}

func maxNum(nums ...int) int {
	if len(nums) == 0 {
		panic("no numbers found")
	}
	max := nums[0]
	for _, n := range nums[1:] {
		if n > max {
			max = n
		}
	}
	return max
}

func main() {
	a := []int{5, 10, 15}
	fmt.Println("sum : ", addAll(a...))
	fmt.Println("average : ", avgAll(a...))
	fmt.Println("max : ", maxNum(a...))
	fmt.Println("min : ", maxNum(a...))
}
