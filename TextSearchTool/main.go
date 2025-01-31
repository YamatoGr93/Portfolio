package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func searchInFile(filePath string, searchTerm string) {
	file, _ := os.Open(filePath)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), searchTerm) {
			fmt.Println(scanner.Text())
		}
	}
}

func main() {
	searchInFile("example.txt", "search_term")
}
