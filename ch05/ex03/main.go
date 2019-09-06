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
		case "script", "style":
			// skip
		default:
			// triming
			text := strings.Trim(n.Data, "\t\n")
			if text != "" {
				content = append(content, text)
			}
		}
	}

	if n.FirstChild != nil {
		content = GetTextContent(content, n.FirstChild)
	}

	if n.NextSibling != nil {
		return GetTextContent(content, n.NextSibling)
	}

	return content
}
