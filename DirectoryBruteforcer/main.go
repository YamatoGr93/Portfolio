package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strings"
	"sync"
)

// scanDirectory tests a single directory path
func scanDirectory(url, path string, statusFilter int, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()
	fullURL := strings.TrimRight(url, "/") + "/" + path

	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Printf("[ERROR] Failed to request: %s\n", fullURL)
		return
	}
	defer resp.Body.Close()

	if statusFilter == 0 || resp.StatusCode == statusFilter {
		results <- fmt.Sprintf("[FOUND] %s (%d)", fullURL, resp.StatusCode)
	}
}

// readWordlist reads a wordlist file into a slice
func readWordlist(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open wordlist: %w", err)
	}
	defer file.Close()

	var words []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}
	return words, scanner.Err()
}

func main() {
	// Command-line flags
	wordlistPath := flag.String("w", "wordlists/common.txt", "Path to the wordlist")
	threadCount := flag.Int("t", 10, "Number of concurrent threads")
	statusFilter := flag.Int("s", 0, "Filter results by HTTP status code (0 = no filter)")
	outputFilePath := flag.String("o", "results.txt", "File to save the results")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: go run main.go <url> [flags]")
		os.Exit(1)
	}

	targetURL := flag.Arg(0)

	// Read wordlist
	wordlist, err := readWordlist(*wordlistPath)
	if err != nil {
		fmt.Printf("Error reading wordlist: %v\n", err)
		os.Exit(1)
	}

	// Open output file
	outputFile, err := os.Create(*outputFilePath)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)

	// Set up concurrency
	var wg sync.WaitGroup
	results := make(chan string)

	// Start a Goroutine to write results to the file
	go func() {
		for result := range results {
			fmt.Println(result) // Still prints to terminal for visibility
			writer.WriteString(result + "\n")
		}
	}()

	// Launch Goroutines for directory scanning
	sem := make(chan struct{}, *threadCount)
	for _, path := range wordlist {
		wg.Add(1)
		sem <- struct{}{}
		go func(path string) {
			defer func() { <-sem }()
			scanDirectory(targetURL, path, *statusFilter, &wg, results)
		}(path)
	}

	wg.Wait()
	close(results)

	// Flush results to file
	writer.Flush()
	fmt.Printf("\nResults have been saved to %s\n", *outputFilePath)
}
