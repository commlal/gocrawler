package main

import (
	"os"
	"fmt"
)

func main () {
    argsGiven := os.Args[1:]
	if len(argsGiven) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(argsGiven) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		pages := make(map[string]int)
		crawlPage(argsGiven[0], argsGiven[0], pages)
		for nURL, count := range pages {
			fmt.Printf("%d - %s\n", count, nURL)
		}
	}
}
