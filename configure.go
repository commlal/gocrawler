package main

import (
	"fmt"
	"net/url"
	"sync"
)

type config struct {
	pages		map[string]PageData
	baseURL		*url.URL
	mu			*sync.Mutex
	cc			chan struct{}
	wg			*sync.WaitGroup
	maxPages	int
}

func (cfg *config) addPageVisit(normalizedURL string) (shouldVisit bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()

	if _, visited := cfg.pages[normalizedURL]; visited {
		return false
	}

	cfg.pages[normalizedURL] = PageData{URL: normalizedURL}
	return true
}

func (cfg *config) setPageData(normalizedURL string, data PageData) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	cfg.pages[normalizedURL] = data
}

func configure(rawBaseURL string, maxConcurrency int, maxPages int) (*config, error) {
	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		return nil, fmt.Errorf("couldn't parse base URL: %v", err)
	}
	return &config{
		pages:		make(map[string]PageData),
		baseURL:	baseURL,
		mu:			&sync.Mutex{},
		cc:			make(chan struct{}, maxConcurrency),
		wg:			&sync.WaitGroup{},
		maxPages:	maxPages,
	}, nil
}
