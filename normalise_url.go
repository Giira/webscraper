package main

import (
	"fmt"
	"strings"
)

func normaliseURL(url string) (string, error) {
	tmp, err := removePrefix(url)
	if err != nil {
		return "", err
	}
	out, _ := strings.CutSuffix(tmp, "/")
	return out, nil
}

func removePrefix(url string) (string, error) {
	if strings.HasPrefix(url, "https://") {
		out, ok := strings.CutPrefix(url, "https://")
		if !ok {
			return "", fmt.Errorf("Failed to remove prefix")
		}
		return out, nil
	} else if strings.HasPrefix(url, "http://") {
		out, ok := strings.CutPrefix(url, "http://")
		if !ok {
			return "", fmt.Errorf("Failed to remove prefix")
		}
		return out, nil
	}
	return url, nil
}
