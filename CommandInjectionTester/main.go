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

// scanPath tests a single directory or path for accessibility
func scanPath(url, path string, wg *sync.WaitGroup, results chan<- string) {
	defer wg.Done()

	// Construct the full URL
	fullURL := strings.TrimRight(url, "/") + "/" + strings.TrimLeft(path, "/")

	// Send HTTP request
	resp, err := http.Get(fullURL)
	if err != nil {
		fmt.Printf("[ERROR] Failed to request: %s\n", fullURL)
		return
	}
	defer resp.Body.Close()

	// Save successful paths
	if resp.StatusCode == 200 {
		results <- fmt.Sprintf("[FOUND] %s (%d)", fullURL, resp.StatusCode)
	}
}

// readWordlist reads paths from a given wordlist file
func readWordlist(filePath string) ([]string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open wordlist: %w", err)
	}
	defer file.Close()

	var paths []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		paths = append(paths, scanner.Text())
	}
	return paths, scanner.Err()
}

func main() {
	// Command-line flags
	var wordlist string
	var output string
	var threads int

	flag.StringVar(&wordlist, "w", "wordlists/common.txt", "Path to the wordlist")
	flag.StringVar(&output, "o", "results.txt", "File to save the results")
	flag.IntVar(&threads, "t", 10, "Number of concurrent threads")
	flag.Parse()

	// Validate URL input
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go <URL> [-w wordlist] [-o output] [-t threads]")
		os.Exit(1)
	}
	url := flag.Arg(0)

	// Read paths from the wordlist
	paths, err := readWordlist(wordlist)
	if err != nil {
		fmt.Printf("Error reading wordlist: %v\n", err)
		os.Exit(1)
	}

	// Open output file for saving results
	outputFile, err := os.Create(output)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outputFile.Close()

	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	// Set up concurrency
	var wg sync.WaitGroup
	results := make(chan string)

	// Goroutine for saving results
	go func() {
		for result := range results {
			fmt.Println(result)               // Output to console
			writer.WriteString(result + "\n") // Save to file
		}
	}()

	// Semaphore for limiting concurrency
	sem := make(chan struct{}, threads)
	for _, path := range paths {
		wg.Add(1)
		sem <- struct{}{}
		go func(p string) {
			defer func() { <-sem }()
			scanPath(url, p, &wg, results)
		}(path)
	}

	// Wait for all goroutines to finish
	wg.Wait()
	close(results)

	fmt.Printf("\nResults have been saved to %s\n", output)
}
