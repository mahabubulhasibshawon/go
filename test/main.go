package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i:= 0 ; i < t; i++ {
		var n int
		fmt.Scan(&n)
	req_berry := n * 3
	fmt.Println((req_berry * 4)/6)
	}
}
