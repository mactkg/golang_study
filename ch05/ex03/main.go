package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	res := GetTextContent(nil, doc)
	fmt.Printf(strings.Join(res, "\n"))
}

func GetTextContent(content []string, n *html.Node) []string {
	if n.Type == html.TextNode {
		switch n.Parent.Data {
		case "html", "body", "script", "style":
			goto Traverse
		}

		// triming
		text := strings.TrimLeft(n.Data, "\t\n")
		if text != "" {
			content = append(content)
		}
	}

Traverse:
	if n.FirstChild != nil {
		content = GetTextContent(content, n.FirstChild)
	}

	if n.NextSibling != nil {
		return GetTextContent(content, n.NextSibling)
	}

	return content
}
