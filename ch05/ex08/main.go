package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Printf("Usage: go run main.go [url] [id]")
		os.Exit(1)
	}

	findElementByIDInURL(os.Args[1], os.Args[2])
}

func findElementByIDInURL(url string, id string) (err error) {
	resp, err := http.Get(url)
	if err != nil {
		err = fmt.Errorf("Error: something wrong...(url: %v, error: %v)", url, err)
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("Error at parsing HTML: %s", err)
		return
	}

	node := ElementByID(doc, id)
	if node != nil {
		fmt.Printf("Found! %v\n", node)
	} else {
		fmt.Printf("Not found...\n")
	}
	return nil
}

func ElementByID(doc *html.Node, id string) *html.Node {
	pre := func(n *html.Node) (*html.Node, bool) {
		if n.Type == html.ElementNode {
			for _, v := range n.Attr {
				if v.Key == "id" && v.Val == id {
					return n, false
				}
			}
		}
		return nil, true
	}

	node, _ := forEachNode(doc, pre, nil)
	return node
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) (*html.Node, bool)) (*html.Node, bool) {
	if pre != nil {
		node, ok := pre(n)
		if !ok {
			return node, ok
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		node, ok := forEachNode(c, pre, post)
		if !ok {
			return node, ok
		}
	}

	if post != nil {
		node, ok := post(n)
		if !ok {
			return node, ok
		}
	}

	return nil, true
}
