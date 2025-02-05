package main

import "fmt"

var a = 10

func main() {
	fmt.Println("main function call")
	fmt.Println(a)
}

func init() {
	fmt.Println("init functions call")
	fmt.Println(a)
	fmt.Println("variable declare in init function")
	a = 20
	fmt.Println(a)
}
