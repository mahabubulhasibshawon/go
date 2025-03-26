package main

import "fmt"

func modifyArray(arr *[5]int, index int, newValue int) {
	// আপনার কোড এখানে
	fmt.Println(arr[index])
	arr[index] = newValue
}

func main() {
	arr := [5]int{1, 2, 3, 4, 5}
	modifyArray(&arr, 2, 10)
	fmt.Println(arr) // আউটপুট: [1 2 10 4 5]
}
