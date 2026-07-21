package main

import (
	"fmt"
	"os"
)

func main() {
	argswithP := os.Args
	args := argswithP[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %v\n", args[0])
		pages := make(map[string]int)
		crawlPage(args[0], args[0], pages)
		fmt.Print("Pages crawled:\nPage: Visits\n")
		for key, value := range pages {
			fmt.Printf("%v: %v\n", key, value)
		}
	}
}
