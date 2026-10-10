package main

import (
	"net/url"
	"fmt"
	"strings"
)

func normalizeURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("ERROR -- Could not parse URL: %w", err)
	}
	fullPath := strings.ToLower(strings.TrimSuffix(parsedURL.Host + parsedURL.Path, "/"))
	return fullPath, nil
}