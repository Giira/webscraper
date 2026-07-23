package main

import (
	"fmt"
	"net/url"
	"strings"
	"sync"
)

type config struct {
	pages              map[string]PageData
	baseURL            *url.URL
	mu                 *sync.Mutex
	concurrencyControl chan struct{}
	wg                 *sync.WaitGroup
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.wg.Add(1)
	if !strings.HasPrefix(rawCurrentURL, cfg.baseURL.String()) {
		return
	}

	normURL, err := normaliseURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error in normaliseURL: %v", err)
		return
	}

	if !cfg.addPageVisit(normURL) {
		return
	}

	fmt.Printf("Crawling %v\n", normURL)
	body, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error in getHTML: %v\n", err)
		return
	}

	pd := extractPageData(body, normURL)
	pd.Visits = 1
	cfg.mu.Lock()
	cfg.pages[normURL] = pd
	cfg.mu.Unlock()
	for _, url := range pd.OutgoingLinks {
		go cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalisedURL string) (isFirst bool) {
	cfg.mu.Lock()
	pd, ok := cfg.pages[normalisedURL]
	if !ok {
		return true
	}
	pd.Visits++
	cfg.pages[normalisedURL] = pd
	cfg.mu.Unlock()
	return false
}
