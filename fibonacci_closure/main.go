package main

import "fmt"

func fibonacci() func() int {
	num1 := 0
	num2 := 1

	return func() int {
		res1 := num1 + num2
		res2 := res1 + num2
		return res2
	}
}

func main() {
	fib := fibonacci()
	fmt.Println(fib)
	fmt.Println(fib)
	fmt.Println(fib)

}

func init() {
	fmt.Println("==== Fibonacci Closure ====")
}
