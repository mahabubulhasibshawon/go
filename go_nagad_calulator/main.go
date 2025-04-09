package main

import "fmt"

func main() {
	var cout, f_main, f_res float32
	for j := 3600; j < 3900; j++ {
		fmt.Println(j)
		cout = j * .125
		f_main = j - int(cout)
	}
}
