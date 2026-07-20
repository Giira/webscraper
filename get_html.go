package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func getHTML(rawURL string) (string, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", rawURL, nil)
	if err != nil {
		return "", fmt.Errorf("error making new request: %v", err)
	}
	req.Header.Set("User-Agent", "GoCrawler/1.0")

	res, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("error making http request: %v", err)
	}
	defer res.Body.Close()

	if res.StatusCode >= 400 {
		return "", fmt.Errorf("Failed to get proper response: error code %v: %v", res.StatusCode, err)
	}

	conType := res.Header.Get("Content-Type")
	if !strings.HasPrefix(conType, "text/html") {
		return "", fmt.Errorf("error: response is not html: content-type: %v: %v", conType, err)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", fmt.Errorf("error reasing response body: %v", err)
	}
	return string(body), nil
}
