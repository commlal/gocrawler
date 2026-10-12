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
		fmt.Printf("starting crawl of: %s\n", argsGiven[0])
		os.Exit(0)
	}
}
