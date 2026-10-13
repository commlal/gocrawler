package main

import (
	"os"
	"fmt"
	"github.com/fatih/color"
)

func main () {
    argsGiven := os.Args[1:]
	if len(argsGiven) < 1 {
		fmt.Println("No website provided")
		os.Exit(1)
	} else if len(argsGiven) > 1 {
		fmt.Println("Too many arguments provided")
		os.Exit(1)
	} 
	rawBaseURL := argsGiven[0]
	maxConcurrency := 1
	cfg, err := configure(rawBaseURL, maxConcurrency)
	if err != nil {
		fmt.Printf("ERROR -- Could not create cfg: %v", err)
		return
	}

	fmt.Printf(color.GreenString("GOCRAWLER -- Begining crawl of %s\n", rawBaseURL))

	cfg.wg.Add(1)
	go cfg.crawlPage(rawBaseURL)
	cfg.wg.Wait()

	fmt.Println(color.GreenString("GOCRAWLER -- Crawling Complete"))
}
