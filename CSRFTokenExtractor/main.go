package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
)

// ExtractCSRFToken extracts CSRF tokens from a response body using a regular expression
func ExtractCSRFToken(body string) []string {
	// Regex pattern to match CSRF tokens
	tokenRegex := regexp.MustCompile(`(?i)(name=["']csrf-token["']|value=["'](.*?csrf.*?)["'])`)
	matches := tokenRegex.FindAllStringSubmatch(body, -1)

	var tokens []string
	for _, match := range matches {
		if len(match) > 2 {
			tokens = append(tokens, match[2])
		}
	}
	return tokens
}

func main() {
	// Command-line flags
	url := flag.String("url", "", "Target URL to extract CSRF tokens from")
	output := flag.String("o", "csrf_tokens.txt", "File to save extracted CSRF tokens")
	flag.Parse()

	// Check if URL is provided
	if *url == "" {
		fmt.Println("Error: URL is required. Use -url flag to specify the target.")
		os.Exit(1)
	}

	// Make an HTTP GET request to the URL
	resp, err := http.Get(*url)
	if err != nil {
		fmt.Printf("Error making GET request to %s: %v\n", *url, err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		os.Exit(1)
	}
	body := string(bodyBytes)

	// Extract CSRF tokens
	tokens := ExtractCSRFToken(body)
	if len(tokens) == 0 {
		fmt.Println("No CSRF tokens found.")
		return
	}

	// Print tokens to the terminal
	fmt.Println("Extracted CSRF Tokens:")
	for _, token := range tokens {
		fmt.Println(token)
	}

	// Save tokens to a file
	file, err := os.Create(*output)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	for _, token := range tokens {
		_, _ = file.WriteString(token + "\n")
	}

	fmt.Printf("CSRF tokens have been saved to %s\n", *output)
}
