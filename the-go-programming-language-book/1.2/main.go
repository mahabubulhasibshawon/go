package main

import (
	"fmt"
	"os"
	"strings"
)

// func main() {
// 	var s, sep string
// 	for i := 0; i < len(os.Args); i++ {
// 		s += sep + os.Args[i]
// 		sep = " "
// 	}
// 	fmt.Println(s)
// }

// func main() {
// 	s, sep := "", ""
// 	for index, arg := range os.Args {
// 		s += sep + arg
// 		sep = " "
// 		fmt.Println(index, arg)
// 	}
// 	// fmt.Println(s)
// }

func main() {
 fmt.Println(strings.Join(os.Args[1:], " "))
 }