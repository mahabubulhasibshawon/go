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
	slice := numbers[1:4]
	fmt.Println("slice example : ", slice)
	fmt.Println("length : ", len(numbers))
	fmt.Println("capacity : ", cap(numbers))
	fmt.Printf("numbers has a data type of : %T\n", numbers)

	// slice using make
	s := make([]int, 30, 50)
	fmt.Println("slice using make : ", s)
	fmt.Println("Length : ", len(s))
	fmt.Println("capacity ", cap(s))
	// modifing this slice
	s[1] = 11
	fmt.Println("after modifing : ", s)

	// emply slice
	var empty_slice []int
	fmt.Println("Empty slice : ", empty_slice)
	empty_slice = append(empty_slice, 1, 3, 2)
	fmt.Println("after add  element : ", empty_slice)

	// slice underlying array
	// rule => 1024 -> double : (>1024 -> 25%)
	var x []int
	x = append(x, 1)
	x = append(x, 2)
	x = append(x, 3)
	fmt.Println(x)
	y := x

	fmt.Println(y)
	x = append(x, 4)
	fmt.Println(x)
	y = append(y, 5)

	x[0] = 10

	fmt.Println(x)
	fmt.Println(y)
}
