package main

import (
	"fmt"
	"os"
)

func main() {
	argswithP := os.Args
	args := argswithP[1:]

	if len(args) < 1 {
		fmt.Println("no website provided")
		os.Exit(1)
	} else if len(args) > 1 {
		fmt.Println("too many arguments provided")
		os.Exit(1)
	} else {
		fmt.Printf("starting crawl of: %v\n", args[0])
		html, err := getHTML(args[0])
		if err != nil {
			fmt.Print(err)
			os.Exit(1)
		}
		fmt.Print(html)
	}

}
