package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	var out []string
	out = strings.Split(html, "<h1>")
	if len(out) == 1 {
		out = strings.Split(out[0], "<h2>")
		if len(out) == 1 {
			return ""
		}
		out = strings.Split(out[1], "</h2>")
	}
	out = strings.Split(out[1], "</h1>")

	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatal("failed to prepare for goquery search: %v", err)
	}
	if doc.Find("<h1>") {

	}
}
