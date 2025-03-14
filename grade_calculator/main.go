package main

import "fmt"

func gradeCalculator(score int) {
	if score > 89 && score < 101 {
		fmt.Println("A")
	} else if score >= 80 && score <= 89 {
		fmt.Println("B")
	} else if score >= 70 && score <= 79 {
		fmt.Println("C")
	}
}

func main() {
	var score int
	fmt.Print("Enter the score: ")
	fmt.Scanln(&score)
	gradeCalculator(score)
}
