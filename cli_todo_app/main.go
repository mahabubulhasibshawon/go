package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Task struct {
	ID        int
	Title     string
	Deadline  string
	Completed bool
}

func taskManager() func(string, string) {
	var tasks []Task
	var idCounter int

	return func(command string, input string) {
		switch command {
		case "add":
			// Split input as: "title | deadline"
			parts := strings.SplitN(input, "|", 2)
			if len(parts) < 2 {
				fmt.Println("Usage: add <title> | <deadline>")
				return
			}
			title := strings.TrimSpace(parts[0])
			deadline := strings.TrimSpace(parts[1])
			idCounter++
			tasks = append(tasks, Task{
				ID:       idCounter,
				Title:    title,
				Deadline: deadline,
			})
			fmt.Println("Added task:", title)

		case "list":
			if len(tasks) == 0 {
				fmt.Println("No tasks found.")
				return
			}
			for _, t := range tasks {
				status := " "
				if t.Completed {
					status = "x"
				}
				fmt.Printf("[%s] %d: %s (Deadline: %s)\n", status, t.ID, t.Title, t.Deadline)
			}

		case "delete":
			id, _ := strconv.Atoi(input)
			for i, t := range tasks {
				if t.ID == id {
					tasks = append(tasks[:i], tasks[i+1:]...)
					fmt.Println("Deleted task ID:", id)
					return
				}
			}
			fmt.Println("Task not found.")

		case "complete":
			id, _ := strconv.Atoi(input)
			for i := range tasks {
				if tasks[i].ID == id {
					tasks[i].Completed = true
					fmt.Println("Completed task ID:", id)
					return
				}
			}
			fmt.Println("Task not found.")
		}
	}
}

func main() {
	manager := taskManager()
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to Go To-Do CLI")
	fmt.Println("Commands:")
	fmt.Println("- add <title> | <deadline>")
	fmt.Println("- delete <id>")
	fmt.Println("- complete <id>")
	fmt.Println("- list")
	fmt.Println("- exit")

	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		if line == "exit" {
			break
		}

		parts := strings.SplitN(line, " ", 2)
		command := parts[0]
		var arg string
		if len(parts) > 1 {
			arg = parts[1]
		}

		manager(command, arg)
	}
}
