package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
)

func main() {
	// Command-line flags
	wordlist := flag.String("w", "wordlists/common.txt", "Path to the wordlist file")
	outputFile := flag.String("o", "results.txt", "File to save found subdomains")
	silent := flag.Bool("s", false, "Enable silent mode (no console output)")
	threads := flag.Int("t", 10, "Number of threads for parallel lookups")
	flag.Parse()

	// Ensure a domain is provided
	if len(flag.Args()) < 1 {
		fmt.Println("Usage: go run main.go <domain> [-w wordlist] [-o output_file] [-s] [-t threads]")
		os.Exit(1)
	}
	domain := flag.Args()[0]

	// Open the wordlist file
	file, err := os.Open(*wordlist)
	if err != nil {
		fmt.Println("Error opening wordlist:", err)
		os.Exit(1)
	}
	defer file.Close()

	// Open the output file for writing
	out, err := os.Create(*outputFile)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		os.Exit(1)
	}
	defer out.Close()

	writer := bufio.NewWriter(out)
	defer writer.Flush()

	// Set up concurrency controls
	var wg sync.WaitGroup
	sem := make(chan struct{}, *threads) // Semaphore for limiting Goroutines

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		wg.Add(1)
		sem <- struct{}{} // Block if max threads are reached
		go func(subdomain string) {
			defer wg.Done()
			defer func() { <-sem }() // Release semaphore
			_, err := net.LookupHost(subdomain)
			if err == nil {
				if !*silent {
					fmt.Println("[+] Found subdomain:", subdomain)
				}
				_, _ = writer.WriteString(subdomain + "\n")
			}
		}(scanner.Text() + "." + domain)
	}

	wg.Wait() // Wait for all Goroutines to finish

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading wordlist:", err)
	}
}
