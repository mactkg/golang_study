package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	f, err := os.Open("./test.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	doc, err := html.Parse(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	images := ElementByTagName(doc, "img")
	headings := ElementByTagName(doc, "h1", "h2", "h3")

	fmt.Println("images")
	for _, v := range images {
		fmt.Printf("%v: %v\n", v.Data, v.Attr)
	}

	fmt.Println("headings")
	for _, v := range headings {
		fmt.Printf("%v: %v\n", v.Data, v.Attr)
	}
}

func ElementByTagName(doc *html.Node, name ...string) []*html.Node {
	res := []*html.Node{}
	if len(name) == 0 {
		return res
	}

	forEachNode(doc, func(n *html.Node) {
		if n.Type == html.ElementNode {
			for _, v := range name {
				if n.Data == v {
					res = append(res, n)
				}
			}
		}
	}, nil)
	return res
}

func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}

	if post != nil {
		post(n)
	}
}
