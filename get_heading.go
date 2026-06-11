package main

import "strings"

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
	return out[0]

	doc := strings.NewReader(html)

	if doc.Find("<h1>") {

	}
}
