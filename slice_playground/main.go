package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var slice []int

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n==== Slice Playground ====")
		fmt.Println("Current Slice:", slice)
		fmt.Println("1. Append element")
		fmt.Println("2. Remove element by index")
		fmt.Println("3. Reverse slice")
		fmt.Println("4. Search for an element")
		fmt.Println("5. Exit")
		fmt.Print("Choose an option: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		switch input {
		case "1":
			appendElement(reader)
		case "2":
			removeElement(reader)
		case "3":
			reverseSlice()
		case "4":
			searchElement(reader)
		case "5":
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option. Try again.")
		}
	}
}

func appendElement(reader *bufio.Reader) {
	fmt.Print("Enter number to append: ")
	input, _ := reader.ReadString('\n')
	num, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}
	slice = append(slice, num)
	fmt.Println("Element appended.")
}

func removeElement(reader *bufio.Reader) {
	fmt.Print("Enter index to remove: ")
	input, _ := reader.ReadString('\n')
	index, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil || index < 0 || index >= len(slice) {
		fmt.Println("Invalid index.")
		return
	}
	slice = append(slice[:index], slice[index+1:]...)
	fmt.Println("Element removed.")
}

func reverseSlice() {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
	fmt.Println("Slice reversed.")
}

func searchElement(reader *bufio.Reader) {
	fmt.Print("Enter element to search: ")
	input, _ := reader.ReadString('\n')
	target, err := strconv.Atoi(strings.TrimSpace(input))
	if err != nil {
		fmt.Println("Invalid input.")
		return
	}
	found := false
	for i, val := range slice {
		if val == target {
			fmt.Printf("Found %d at index %d\n", target, i)
			found = true
		}
	}
	if !found {
		fmt.Println("Element not found.")
	}
}
