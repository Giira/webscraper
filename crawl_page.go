package main

import (
	"fmt"
	"net/url"
	"strings"
)

func crawlPage(rawBaseURL string, rawCurrentURL string, pages map[string]int) {
	if !strings.HasPrefix(rawCurrentURL, rawBaseURL) {
		return
	}

	normURL, err := normaliseURL(rawCurrentURL)
	if err != nil {
		fmt.Printf("error in normaliseURL: %v", err)
		return
	}

	i, ok := pages[normURL]
	if !ok {
		pages[normURL] = 1
	} else {
		pages[normURL] = i + 1
		return
	}

	fmt.Printf("Crawling %v\n", normURL)
	body, err := getHTML(rawCurrentURL)
	if err != nil {
		fmt.Printf("error in getHTML: %v\n", err)
		return
	}

	baseURL, err := url.Parse(rawBaseURL)
	if err != nil {
		fmt.Printf("failed to parse url: %v\n", err)
		return
	}
	urls, err := getURLsFromHTML(body, baseURL)
	for _, url := range urls {
		crawlPage(rawBaseURL, url, pages)
	}
}
