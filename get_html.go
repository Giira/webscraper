package main

import (
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func getHeadingFromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("failed to prepare for goquery search: %v", err)
	}

	selection := doc.Find("h1").First()
	heading := selection.Text()
	if heading != "" {
		return heading
	}

	selection = doc.Find("h2").First()
	heading = selection.Text()
	if heading != "" {
		return heading
	}

	return ""
}

func getFirstParagraphFromHTML(html string) string {
	reader := strings.NewReader(html)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("failed to prepare for goquery search: %v", err)
	}

	selection := doc.Find("p").First()
	paragraph := selection.Text()
	return paragraph
}
