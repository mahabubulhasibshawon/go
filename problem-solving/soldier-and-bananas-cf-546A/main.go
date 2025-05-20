package main

import "fmt"

func main () {
	var k, n, w int
	fmt.Scan(&k, &n, &w)
	total_cost := k*(w*(w+1))/2
	borrow := total_cost - n
	if borrow < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(borrow)
	}
}