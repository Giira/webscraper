package main

import (
	"fmt"
	"net/url"
	"os"
	"sync"
)

func main() {
	argswithP := os.Args
	args := argswithP[1:]

	maxConcurrency := 5

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %v\n", args[0])

		URL, err := url.Parse(args[0])
		if err != nil {
			fmt.Printf("failed to parse url: %v", err)
			os.Exit(1)
		}

		cfg := &config{
			pages:              make(map[string]PageData),
			baseURL:            URL,
			mu:                 &sync.Mutex{},
			concurrencyControl: make(chan struct{}, maxConcurrency),
			wg:                 &sync.WaitGroup{},
		}

		cfg.wg.Add(1)
		cfg.crawlPage(args[0])
		cfg.wg.Wait()

		fmt.Print("Pages crawled:\nPage: Visits\n")
		for key, pd := range cfg.pages {
			fmt.Printf("%v: %v\n", key, pd.Visits)
		}
	}
}
