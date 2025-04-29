package main

import "fmt"

var a = 10 // data  segmetn

func main() {
	fmt.Println(a)
}

func init() {
	fmt.Println("init")
	a = 20
	fmt.Println(a)
}
