package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Sum(x int) int {
	sum := 0
	for i := 1; i <= x; i++ {
		for j := 0; j < 1000000; j++ {
		}
		time.Sleep(100 * time.Millisecond)
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(3)))
		sum += i
	}
	return sum
}

func main() {
	x := 100

	start := time.Now()
	result := Sum(x)
	elapsed := time.Since(start)

	fmt.Println("Sum of first", x, "numbers:", result)
	fmt.Println("Execution Time:", elapsed)
}
