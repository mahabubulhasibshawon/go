package main

import (
	"fmt"
	"strings"
	"time"
)

func inefficientConcat(args []string) string {
	s, sep := "", ""
	for _, arg := range args {
		s += sep + arg
		sep = " "
	}
	return s
}

func efficientConcat(args []string) string {
	return strings.Join(args, " ")
}

func main() {
	// Simulate command-line arguments or use os.Args[1:]
	args := make([]string, 10000)
	for i := 0; i < len(args); i++ {
		args[i] = "word"
	}

	// Measure inefficient version
	start := time.Now()
	_ = inefficientConcat(args)
	fmt.Println("Inefficient version took:", time.Since(start))

	// Measure efficient version
	start = time.Now()
	_ = efficientConcat(args)
	fmt.Println("Efficient (strings.Join) version took:", time.Since(start))
}
