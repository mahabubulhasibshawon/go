package main

import "fmt"

func processOperation(a int, b int, op func(p int, q int)) {
	op(a, b)
}

func call() func(x int, y int) {
	return substract
}

func mul() func(p, q int) {
	return multiply
}

func add(a int, b int) {
	z := a + b
	fmt.Println(z)
}

func substract(a, b int) {
	z := a - b
	fmt.Println(z)
}

func multiply(a, b int) {
	fmt.Println(a * b)
}

func main() {

	// fmt.Println(add(5, 5))
	processOperation(4, 5, add)
	sum := call()
	sum(4, 9)
	mult := mul()
	mult(9, 9)
}
