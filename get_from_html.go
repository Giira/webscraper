package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type PageData struct {
	URL            string   `json:"url"`
	Heading        string   `json:"heading"`
	FirstParagraph string   `json:"first_paragraph"`
	OutgoingLinks  []string `json:"outgoing_links"`
	ImageURLs      []string `json:"image_urls"`
	Visits         int      `json:"visits"`
}

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

	main := doc.Find("main")
	var selection *goquery.Selection
	if main.Length() > 0 {
		selection = main.Find("p").First()
	} else {
		selection = doc.Find("p").First()
	}
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
	doc.Find("a[href]").Each(func(_ int, s *goquery.Selection) {
		urlString, ok := s.Attr("href")
		if !ok {
			return
		}

		urlString = strings.TrimSpace(urlString)

		url, err := url.Parse(urlString)
		if err != nil {
			log.Fatalf("error parsing url: %v", err)
		}

		absolute := baseURL.ResolveReference(url)
		out = append(out, absolute.String())
	})
	return out, nil
}

func getImagesFromHTML(htmlBody string, baseURL *url.URL) ([]string, error) {
	reader := strings.NewReader(htmlBody)
	doc, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		log.Fatalf("failed to prepare for goquery search: %v", err)
	}

	var out []string
	doc.Find("img").Each(func(_ int, s *goquery.Selection) {
		imgString, ok := s.Attr("src")
		if !ok {
			return
		}

		img, err := url.Parse(imgString)
		if err != nil {
			log.Fatalf("error parsing url: %v", err)
		}

		absolute := baseURL.ResolveReference(img)
		out = append(out, absolute.String())
	})
	return out, nil
}

func extractPageData(html, pageURL string) PageData {
	page, err := url.Parse(pageURL)
	if err != nil {
		log.Fatalf("error: failed to parse url: %v", err)
	}

	links, err := getURLsFromHTML(html, page)
	if err != nil {
		log.Fatalf("error: failed to get urls: %v", err)
	}

	images, err := getImagesFromHTML(html, page)
	if err != nil {
		log.Fatalf("error: failed to get images: %v", err)
	}

	out := PageData{
		URL:            pageURL,
		Heading:        getHeadingFromHTML(html),
		FirstParagraph: getFirstParagraphFromHTML(html),
		OutgoingLinks:  links,
		ImageURLs:      images,
	}

	return out
}
