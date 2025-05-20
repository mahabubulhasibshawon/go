package main

import "fmt"

func main() {
	var X, t int
	fmt.Scan(&X, &t)
	if X < t {
		fmt.Println(0)
	} else {
		fmt.Println(X-t)
	}
}