package main

import "fmt"

func swap(a *int, b *int) {
	// আপনার কোড এখানে
	*a = 20
	*b = 10
}

func main() {
	x := 10
	y := 20
	swap(&x, &y)
	fmt.Println(x, y) // আউটপুট: 20 10
}
