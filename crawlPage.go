package main

import (
	"log"
	"strings"
)

func crawlPage(rawBaseURL, rawCurrentURL string, pages map[string]int) {
	if !strings.Contains(rawCurrentURL, rawBaseURL) && !strings.HasPrefix(rawCurrentURL, "/") {// Checks if current domain OR relative path
		log.Printf("ERROR -- Out of bounds: %s\n" ,rawCurrentURL)
		return 	
	}
	normalCurrentURL, err := normalizeURL(rawCurrentURL)
	if err != nil {
		log.Printf("ERROR -- Can not normalize URL: %s\n" ,rawCurrentURL)
		return
	}
	_, ok := pages[normalCurrentURL]
	if !ok { // URL not in pages map
		pages[normalCurrentURL] = 1
		html, err := getHTML("http://"+normalCurrentURL)
		if err != nil {
			log.Printf("ERROR -- Unable to get data from URL: %v\n", err)
			return
		}
		pageData := extractPageData(html, "http://"+normalCurrentURL)
		for _, page := range pageData.OutgoingLinks {
			crawlPage(rawBaseURL, page, pages)
		}
	} else {
		pages[normalCurrentURL] = pages[normalCurrentURL] + 1
	} 
}