package main

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestVisit(t *testing.T) {
	f, err := os.Open("./test.html")
	if err != nil {
		t.Fatal(err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		t.Fatal(err)
	}

	expected := []string{
		"https://cdnjs.cloudflare.com/ajax/libs/mini.css/3.0.1/mini-default.min.css",
		"/",
		"/blogs/1",
		"http://golang.org/",
		"https://golang.org/doc/gopher/gophercolor.png",
		"https://blog.golang.org/gopher",
		"https://commondatastorage.googleapis.com/gtv-videos-bucket/sample/BigBuckBunny.mp4",
		"https://interactive-examples.mdn.mozilla.net/",
		"/media/examples/flower.webm",
		"/media/examples/flower.mp4",
		"https://tip.golang.org/",
		"/media/examples/grapefruit-slice-332-332.jpg",
		"https://contact.me/golang",
		"https://www.google-analytics.com/analytics.js",
	}
	results := Visit(nil, doc)
	for i, v := range results {
		if v != expected[i] {
			t.Fatalf("Expected %v, but got %v", expected[i], v)
		}
	}
}

func TestVisitWithEmptyString(t *testing.T) {
	results := Visit(nil, nil)
	if len(results) != 0 {
		t.Fatalf("Expected len(results) is 0, but got %d", len(results))
	}
}
