package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
	"sync"
)

func main() {
	// Command-line flags
	wordlistFile := flag.String("w", "", "Path to wordlist file")
	outputFile := flag.String("o", "results.txt", "File to save results")
	threads := flag.Int("t", 10, "Number of concurrent threads")
	verbose := flag.Bool("v", false, "Enable verbose output")

	flag.Parse()

	// Debugging output for parsed flags
	fmt.Printf("DEBUG: wordlistFile=%s, outputFile=%s, threads=%d, verbose=%v\n", *wordlistFile, *outputFile, *threads, *verbose)

	// Ensure target URL is provided
	if len(flag.Args()) < 1 {
		fmt.Println("Error: Target URL is required.")
		fmt.Println("Usage: go run main.go <target_url> -w <wordlist_file> -o <output_file> -t <threads> -v")
		os.Exit(1)
	}
	targetURL := flag.Args()[0]

	// Ensure wordlist file is provided
	if *wordlistFile == "" {
		fmt.Println("Error: Wordlist file is required. Use the -w flag to specify it.")
		os.Exit(1)
	}

	// Open the wordlist file
	file, err := os.Open(*wordlistFile)
	if err != nil {
		fmt.Printf("Error opening wordlist file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	// Read the wordlist into memory
	var parameters []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		parameters = append(parameters, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading wordlist file: %v\n", err)
		os.Exit(1)
	}

	// Open the output file
	outFile, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer outFile.Close()

	// Create channels for results and synchronization
	results := make(chan string)
	var wg sync.WaitGroup

	// Semaphore to limit concurrency
	semaphore := make(chan struct{}, *threads)

	// Start processing
	for _, param := range parameters {
		wg.Add(1)
		semaphore <- struct{}{} // Acquire a slot

		go func(param string) {
			defer wg.Done()
			defer func() { <-semaphore }() // Release the slot

			url := fmt.Sprintf("%s?%s=test", targetURL, param)
			if *verbose {
				fmt.Printf("Testing: %s\n", url)
			}

			resp, err := http.Get(url)
			if err != nil {
				if *verbose {
					fmt.Printf("Error testing %s: %v\n", url, err)
				}
				return
			}
			defer resp.Body.Close()

			if resp.StatusCode == http.StatusOK {
				result := fmt.Sprintf("[+] %s", url)
				results <- result
				if *verbose {
					fmt.Println(result)
				}
			}
		}(param)
	}

	// Close the results channel when all workers are done
	go func() {
		wg.Wait()
		close(results)
	}()

	// Write results to the output file
	for result := range results {
		_, err := outFile.WriteString(result + "\n")
		if err != nil {
			fmt.Printf("Error writing to output file: %v\n", err)
		}
	}

	fmt.Println("\nProcessing completed! Results saved to:", *outputFile)
}
