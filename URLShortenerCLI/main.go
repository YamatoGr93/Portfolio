package main

import (
	"fmt"
	"io"
	"net/http"
)

func shortenURL(url string) (string, error) {
	resp, err := http.Get("http://tinyurl.com/api-create.php?url=" + url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	shortURL, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(shortURL), nil
}

func main() {
	shortURL, err := shortenURL("https://example.com")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(shortURL)
}
