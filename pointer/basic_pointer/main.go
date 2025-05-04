package main

import "fmt"

func main() {
	x := "value"
	fmt.Println("1st value of x: ", x)
	p := &x                            // address of pointer => &
	*p = "value changed using pointer" // dereference pointer =>
	fmt.Println("Address of p: ", p)

	fmt.Println("changed value of x: ", x)
}
