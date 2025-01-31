package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	responses := map[string]string{
		"hello": "Hi there!",
		"bye":   "Goodbye!",
	}
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("You: ")
		scanner.Scan()
		input := scanner.Text()

		if input == "exit" {
			break
		}

		if response, found := responses[input]; found {
			fmt.Println("Bot:", response)
		} else {
			fmt.Println("Bot: I don't understand.")
		}
	}
}
