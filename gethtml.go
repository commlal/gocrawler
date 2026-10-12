package main

import (
	"net/http"
	"time"
	"io"
	"fmt"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{
	Timeout: time.Second * 10,
	}

	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("Failure to create request: %v", err)
	}
	req.Header.Set("User-Agent", "GoCrawler/0.1")

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Network error: %v", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode > 399 {
		return "", fmt.Errorf("HTTP error: %s", resp.Status)
	}

	contentType := resp.Header.Get("Content-Type")
	if !strings.Contains(contentType, "text/html") {
		return "", fmt.Errorf("Received non-HTML resp: %s", contentType)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Could not read resp body: %v", err)
	}

	return string(body), nil
}