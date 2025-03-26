// package main

// import "fmt"

// func double(ptr *int) {
// 	*ptr = 10
// }

//	func main() {
//		x := 5
//		double(&x)
//		fmt.Println(x)
//	}

package main

import "fmt"

func calculate(a *int, b *int, sum *int, product *int) {
	// আপনার কোড এখানে
	*sum = *a + *b
	*product = *a * *b
}

func main() {
	x := 5
	y := 3
	var s, p int
	calculate(&x, &y, &s, &p)
	fmt.Println("Sum:", s, "Product:", p) // আউটপুট: Sum: 8 Product: 15
}
