package main

import (
	"log"
	"net/url"
	"fmt"
	"github.com/fatih/color"
)

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.mu.Lock()
	pgsCrawled := len(cfg.pages)
	cfg.mu.Unlock()

	if cfg.maxPages <= pgsCrawled {  // Check if more than maxPages reached
		fmt.Printf("INFO -- Page limit reached. %v not crawled\n", rawCurrentURL)
		cfg.wg.Done()
		return
	}

	cfg.cc <- struct{}{} // Concurrancy Control
	defer func() {
		<-cfg.cc
		cfg.wg.Done()
	}()

	currentURL, err := url.Parse(rawCurrentURL)
	if currentURL.Hostname() != cfg.baseURL.Hostname() { //Scope check
		return
	}

	normalCurrentURL, err := normalizeURL(rawCurrentURL) 
	if err != nil {
		log.Printf("ERROR -- Can not normalize URL: %s\n" ,rawCurrentURL)
		return
	}

	visitNeeded := cfg.addPageVisit(normalCurrentURL) // Call to helper method, returns true if not in map
	if !visitNeeded {
		return
	}

	fmt.Printf(color.GreenString("GOCRAWLER -- Crawling: %s\n", normalCurrentURL))
	
	data, err := getHTML("http://"+normalCurrentURL)
	if err != nil {
		fmt.Printf("ERROR -- Unable to fetch HTML data: %v\n", err)
		return
	}

	pageData := extractPageData(data, normalCurrentURL)
	cfg.setPageData(normalCurrentURL, pageData)

	for _, nextURL := range pageData.OutgoingLinks { // New recursive call
		cfg.wg.Add(1)
		go cfg.crawlPage(nextURL)
	}
	}


