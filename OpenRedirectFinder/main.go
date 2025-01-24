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

// Predefined redirect parameters
var redirectParams = []string{"redirect", "url", "next", "target", "dest", "link", "out"}

// Predefined payloads for testing
var payloads = []string{
	"https://evil.com",
	"//evil.com",
	"/\\evil.com",
	"\\evil.com",
	"http://localhost",
}

// Function to check for redirection
func checkRedirect(targetURL, param, payload string) bool {
	testURL := fmt.Sprintf("%s?%s=%s", targetURL, param, payload)
	resp, err := http.Get(testURL)
	if err != nil {
		fmt.Printf("[!] Error sending request to %s: %v\n", testURL, err)
		return false
	}
	defer resp.Body.Close()

	// Check for redirection
	if resp.StatusCode >= 300 && resp.StatusCode < 400 {
		location := resp.Header.Get("Location")
		if strings.Contains(location, payload) {
			return true
		}
	}
	return false
}

// Function to test a single URL with a parameter and payload
func testURL(baseURL, param, payload string) (bool, string) {
	if checkRedirect(baseURL, param, payload) {
		return true, fmt.Sprintf("%s?%s=%s", baseURL, param, payload)
	}
	return false, ""
}

func main() {
	// Command-line flags
	outputFile := flag.String("o", "redirects.txt", "File to save vulnerable URLs")
	wordlistFile := flag.String("w", "", "Path to custom wordlist of redirect parameters")
	threads := flag.Int("t", 10, "Number of concurrent threads")
	verbose := flag.Bool("v", false, "Enable verbose output for debugging")
	flag.Parse()

	// Validate command-line arguments
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Usage: go run main.go <target_url> [flags]")
		os.Exit(1)
	}

	targetURL := args[0]

	// Load custom redirect parameters if provided
	if *wordlistFile != "" {
		file, err := os.Open(*wordlistFile)
		if err != nil {
			fmt.Println("[!] Error opening wordlist file:", err)
			os.Exit(1)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			redirectParams = append(redirectParams, scanner.Text())
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("[!] Error reading wordlist file:", err)
			os.Exit(1)
		}
	}

	// Open the output file
	outFile, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("[!] Error creating output file:", err)
		os.Exit(1)
	}
	defer outFile.Close()

	fmt.Printf("Testing target URL: %s\n", targetURL)
	fmt.Printf("Results will be saved to: %s\n\n", *outputFile)

	// Multithreading setup
	var wg sync.WaitGroup
	semaphore := make(chan struct{}, *threads)
	results := make(chan string, 10)

	// Test each parameter and payload combination
	for _, param := range redirectParams {
		for _, payload := range payloads {
			wg.Add(1)
			semaphore <- struct{}{}

			go func(param, payload string) {
				defer wg.Done()
				defer func() { <-semaphore }()

				vulnerable, testedURL := testURL(targetURL, param, payload)
				if vulnerable {
					result := fmt.Sprintf("[+] Vulnerable: %s", testedURL)
					if *verbose {
						fmt.Println(result)
					}
					results <- result
				}
			}(param, payload)
		}
	}

	// Wait for all tests to complete
	go func() {
		wg.Wait()
		close(results)
	}()

	// Write results to the output file
	for result := range results {
		outFile.WriteString(result + "\n")
	}

	fmt.Println("\nTesting completed!")
}
