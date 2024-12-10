package main

import (
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// FetchWebPageWordCount fetches a webpage and returns a map with the count of each word.
func FetchWebPageWordCount(url string) (map[string]int, error) {
	// Fetch the webpage
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch webpage: %w", err)
	}
	defer resp.Body.Close()

	// Check for successful response
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Read the body content
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	// Convert content to lowercase to make word counting case-insensitive
	content := strings.ToLower(string(body))

	// Use a regular expression to extract words
	re := regexp.MustCompile(`\b\w+\b`)
	words := re.FindAllString(content, -1)

	// Count occurrences of each word
	wordCount := make(map[string]int)
	for _, word := range words {
		wordCount[word]++
	}

	return wordCount, nil
}

func main() {
	url := "https://www.example.com" // Replace with the desired webpage URL

	wordCount, err := FetchWebPageWordCount(url)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}

	fmt.Println("Word count:")
	for word, count := range wordCount {
		fmt.Printf("%s: %d\n", word, count)
	}
}
