package main

import (
	"bufio"
	"fmt"
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

	w, i, err := CountWordsAndImages(os.Args[1])
	if err != nil {
		fmt.Printf("Error: something wrong...(url: %v, error: %v)", os.Args[1], err)
		os.Exit(1)
	}
	fmt.Printf("words: %d, Images: %d\n", w, i)
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n.Type == html.ElementNode && n.Data == "img" {
		for _, a := range n.Attr {
			if a.Key == "src" {
				images++
			}
		}
	} else if n.Type == html.TextNode {
		switch n.Parent.Data {
		case "script", "style":
			goto Traverse
		}

		text := strings.TrimLeft(n.Data, "\t\n")
		reader := strings.NewReader(text)
		scanner := bufio.NewScanner(reader)
		scanner.Split(bufio.ScanWords)

		for scanner.Scan() {
			if err := scanner.Err(); err != nil {
				fmt.Fprintf(os.Stderr, "countWordsAndImages	: %v\n", err)
				return
			}

			words++
		}
	}

Traverse:
	if n.FirstChild != nil {
		w, i := countWordsAndImages(n.FirstChild)
		words += w
		images += i
	}

	if n.NextSibling != nil {
		w, i := countWordsAndImages(n.NextSibling)
		words += w
		images += i
	}

	return
}
