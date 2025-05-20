package main

import "fmt"

func main() {
	var N, T, A int
	fmt.Scan(&N, &T, &A)
	if (N/2)+1 <= T || (N/2)+1 <= A {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
}
