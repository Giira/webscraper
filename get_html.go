package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("failed to prepare for goquery search: %v", err)
	}

	for i := range 6 {
		header := fmt.Sprintf("h%v", i+1)
		selection := doc.Find(header).First()
		heading := selection.Text()
		if heading != "" {
			return heading
		}
	}

	return ""
}

func getFirstParagraphFromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("failed to prepare for goquery search: %v", err)
	}

	selection := doc.Find("main").First().Find("p").First()
	paragraph := selection.Text()
	return paragraph
}

func getURLsFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("failed to prepare for goquery search: %v", err)
	}
	var out []string
	_ = doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		out = append(out, s.Text())
	})
	return out, nil
}
