package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: go run main.go [url]")
		os.Exit(1)
	}

	prettyPrint(os.Args[1], os.Stdout)
}

func prettyPrint(url string, w io.Writer) (err error) {
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

	forEachNode(doc, w, startElement, endElement)
	return nil
}

func forEachNode(n *html.Node, w io.Writer, pre, post func(n *html.Node, w io.Writer)) {
	if pre != nil {
		pre(n, w)
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, w, pre, post)
	}

	if post != nil {
		post(n, w)
	}
}

var depth int

func startElement(n *html.Node, w io.Writer) {
	switch n.Type {
	case html.ElementNode:
		fmt.Fprintf(w, "%*s<%s", depth*2, "", n.Data)
		for _, attr := range n.Attr {
			fmt.Fprintf(w, " %s='%s'", attr.Key, attr.Val)
		}
		if n.FirstChild == nil {
			fmt.Fprintf(w, " />\n")
		} else {
			fmt.Fprintf(w, ">\n")
		}
		depth++
	case html.CommentNode:
		fmt.Fprintf(w, "%*s<!-- %s -->\n", depth*2, "", n.Data)
	case html.TextNode:
		trimmed := strings.Trim(n.Data, "\t\n ")
		if len(trimmed) > 0 {
			fmt.Fprintf(w, "%*s%s\n", depth*2, "", trimmed)
		}
	}
}

func endElement(n *html.Node, w io.Writer) {
	switch n.Type {
	case html.ElementNode:
		depth--
		// Write close tag when this node have child element
		if n.FirstChild != nil {
			fmt.Fprintf(w, "%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}
