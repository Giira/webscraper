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
	maxPages           int
}

func (cfg *config) crawlPage(rawCurrentURL string) {
	cfg.concurrencyControl <- struct{}{}
	defer func() {
		<-cfg.concurrencyControl
		cfg.wg.Done()
	}()

	cfg.mu.Lock()
	if len(cfg.pages) >= cfg.maxPages {
		cfg.mu.Unlock()
		return
	}
	cfg.mu.Unlock()

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

	pd := extractPageData(body, rawCurrentURL)
	pd.Visits = 1
	cfg.mu.Lock()
	cfg.pages[normURL] = pd
	cfg.mu.Unlock()
	for _, url := range pd.OutgoingLinks {
		cfg.wg.Add(1)
		go cfg.crawlPage(url)
	}
}

func (cfg *config) addPageVisit(normalisedURL string) (isFirst bool) {
	cfg.mu.Lock()
	defer cfg.mu.Unlock()
	pd, ok := cfg.pages[normalisedURL]
	if !ok {
		return true
	}
	pd.Visits++
	cfg.pages[normalisedURL] = pd
	return false
}
