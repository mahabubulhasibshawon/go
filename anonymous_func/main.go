package main

import "fmt"

var anonumous = func() {
	fmt.Println("anonuymous func call using varible")
}

var sum = func(n1, n2 int) {
	sum := n1 + n2
	fmt.Println("sum is : ", sum)
}

func main() {
	// anonumous func inside main
	func() {
		fmt.Println("hello world")
		anonumous()
		sum(4, 90)
	}()
}
