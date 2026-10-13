package main

import (
	"os"
	"fmt"
	//"github.com/fatih/color"
	"strconv"
)

func main () {
    argsGiven := os.Args[1:]
	if len(argsGiven) < 3 || len(argsGiven) > 3{
		fmt.Println("Incorrect Syntax!")
		fmt.Println("Usage: ./go crawler URL maxConcurrency maxPages")
		os.Exit(1)
	} 

	rawBaseURL := argsGiven[0]
	maxConcurrency, err := strconv.Atoi(argsGiven[1])
	if err != nil {
		fmt.Println("maxConcurrency must be an integer")
		fmt.Println("Usage: ./go crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}
	maxPages, err := strconv.Atoi(argsGiven[2])
	if err != nil {
		fmt.Println("maxPages must be an integer")
		fmt.Println("Usage: ./go crawler URL maxConcurrency maxPages")
		os.Exit(1)
	}
	cfg, err := configure(rawBaseURL, maxConcurrency, maxPages)
	if err != nil {
		fmt.Printf("ERROR -- Could not create cfg: %v", err)
		os.Exit(1)
	}

	fmt.Printf("GOCRAWLER -- Begining crawl of %s\n", rawBaseURL)

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println("GOCRAWLER -- Crawling Complete")
	fmt.Println("GOCRAWLER -- Writing report to 'report.json'")
	err = writeJSONReport(cfg.pages, "report.json")
	if err != nil {
		fmt.Printf("ERROR -- Could not write JSON to file: %v", err)
		os.Exit(1)
	}
}
