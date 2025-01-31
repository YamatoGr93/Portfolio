package main

import (
	"bufio"
	"fmt"
	"os"
)

var tasks []string

func addTask(task string) {
	tasks = append(tasks, task)
}

func listTasks() {
	for i, task := range tasks {
		fmt.Printf("%d. %s\n", i+1, task)
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter command (add/list/exit): ")
		scanner.Scan()
		command := scanner.Text()

		switch command {
		case "add":
			fmt.Print("Task: ")
			scanner.Scan()
			addTask(scanner.Text())
		case "list":
			listTasks()
		case "exit":
			return
		default:
			fmt.Println("Unknown command.")
		}
	}
}
