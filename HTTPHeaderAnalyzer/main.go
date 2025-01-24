package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/http"
	"os"
)

// HeaderInfo holds information about a header, including its recommendation
type HeaderInfo struct {
	Name           string
	Recommendation string
}

func main() {
	// Command-line flags
	outputFile := flag.String("o", "headers_report.txt", "File to save the analysis report")
	customHeaders := flag.String("h", "", "Path to a file containing custom headers to check")
	flag.Parse()

	// Ensure a URL is provided
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go <URL> [-o output_file] [-h custom_headers_file]")
		os.Exit(1)
	}
	url := flag.Args()[0]

	// Default headers to check
	headers := []HeaderInfo{
		{"Content-Security-Policy", "Add a CSP header to control resources that the browser is allowed to load, mitigating XSS attacks."},
		{"Strict-Transport-Security", "Add an HSTS header to enforce HTTPS and protect against protocol downgrade attacks."},
		{"X-Content-Type-Options", "Add this header to prevent browsers from interpreting files as a different MIME type."},
		{"X-Frame-Options", "Add this header to prevent clickjacking attacks by controlling iframe behavior."},
		{"Referrer-Policy", "Add this header to control how much referrer information is shared with other sites."},
	}

	// Load custom headers if provided
	if *customHeaders != "" {
		headers = loadCustomHeaders(*customHeaders)
	}

	// Send HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Analyze headers and collect results
	missingHeaders := []HeaderInfo{}

	fmt.Printf("Analyzing headers for: %s\n", url)
	for _, header := range headers {
		if _, found := resp.Header[header.Name]; !found {
			fmt.Printf("[-] Missing: %s\n", header.Name)
			missingHeaders = append(missingHeaders, header)
		} else {
			fmt.Printf("[+] Found: %s\n", header.Name)
		}
	}

	// Save the results to the output file
	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, header := range missingHeaders {
		_, _ = writer.WriteString(fmt.Sprintf("Missing: %s\nRecommendation: %s\n\n", header.Name, header.Recommendation))
	}

	fmt.Printf("\nAnalysis complete. Report saved to %s\n", *outputFile)
}

// loadCustomHeaders reads custom headers from a file
func loadCustomHeaders(filePath string) []HeaderInfo {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening custom headers file:", err)
		os.Exit(1)
	}
	defer file.Close()

	var headers []HeaderInfo
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		header := scanner.Text()
		headers = append(headers, HeaderInfo{Name: header, Recommendation: "Custom header. No predefined recommendation."})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading custom headers file:", err)
		os.Exit(1)
	}

	return headers
}
