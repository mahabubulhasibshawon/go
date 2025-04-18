package main

import "fmt"

func modifyArray(array [3]int) {
	array[1] = 5
}

func main() {
	arr := [3]int{1, 3, 5}
	modifyArray(arr)

	fmt.Println(arr)
}
