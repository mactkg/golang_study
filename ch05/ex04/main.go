package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	for _, link := range Visit(nil, doc) {
		fmt.Println(link)
	}
}

func Visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode {
		// pick up valid keys of the node
		keys := []string{}
		switch n.Data {
		case "a", "area", "link", "base":
			keys = append(keys, "href")
		case "audio", "embed", "iframe", "img", "input", "script", "source", "video":
			keys = append(keys, "src")
		}

		// match
		for _, a := range n.Attr {
			for _, k := range keys {
				if a.Key == k {
					links = append(links, a.Val)
				}
			}
		}
	}

	if n.FirstChild != nil {
		links = Visit(links, n.FirstChild)
	}

	if n.NextSibling != nil {
		return Visit(links, n.NextSibling)
	}

	return links
}
