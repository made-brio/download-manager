package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
	"sync"
)

// downloadFile handles the file download and saves it locally
func downloadFile(url string, wg *sync.WaitGroup, resultCh chan<- string) {
	defer wg.Done()

	// Extract file name from URL
	fileName := getFileNameFromURL(url)

	// Create file to save the content
	downloadDir := "./downloads/"
	os.MkdirAll(downloadDir, os.ModePerm) // Ensure the directory exists
	file, err := os.Create(downloadDir + fileName)
	if err != nil {
		resultCh <- fmt.Sprintf("Failed to create file %s: %v", fileName, err)
		return
	}
	defer file.Close()

	// Make HTTP request
	resp, err := http.Get(url)
	if err != nil {
		resultCh <- fmt.Sprintf("Failed to download %s: %v", url, err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		resultCh <- fmt.Sprintf("Failed to download %s: HTTP %d", url, resp.StatusCode)
		return
	}

	// Copy content to file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		resultCh <- fmt.Sprintf("Failed to save %s: %v", fileName, err)
		return
	}

	resultCh <- fmt.Sprintf("Successfully downloaded %s", fileName)
}

// getFileNameFromURL extracts the file name from the URL
func getFileNameFromURL(url string) string {
	segments := strings.Split(url, "/")
	return segments[len(segments)-1]
}

func main() {
	var urls []string
	var wg sync.WaitGroup
	resultCh := make(chan string, 10) // Buffered channel for results

	fmt.Println("Enter the URLs (one per line). Type 'done' to finish:")

	// Read URLs from user input
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := scanner.Text()
		if strings.ToLower(input) == "done" {
			break
		}
		if input != "" {
			urls = append(urls, input)
		}
	}

	if len(urls) == 0 {
		fmt.Println("No URLs provided. Exiting.")
		return
	}

	// Start downloading files concurrently
	for _, url := range urls {
		wg.Add(1)
		go downloadFile(url, &wg, resultCh)
	}

	// Wait for all goroutines to finish
	go func() {
		wg.Wait()
		close(resultCh)
	}()

	// Display progress updates
	for result := range resultCh {
		fmt.Println(result)
	}

	fmt.Println("All downloads complete.")
}
