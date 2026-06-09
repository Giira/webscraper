package main

import "strings"

func getHeadingFromHTML(html string) string {
	var out []string
	out = strings.Split(html, "<h1>")
}
