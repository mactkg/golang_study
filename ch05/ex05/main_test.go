package main

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestFetchingOfCountWordsAndImages(t *testing.T) {
	words, images, err := CountWordsAndImages("https://xkcd.com/1570/")
	if err != nil {
		t.Fatal(err)
	}

	if words == 0 {
		t.Fatalf("Can't count words")
	}
	if images == 0 {
		t.Fatalf("Can't count images")
	}
}

func TestCountWordsAndImages(t *testing.T) {
	f, err := os.Open("./test.html")
	if err != nil {
		t.Fatal(err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		t.Fatal(err)
	}
	words, images := countWordsAndImages(doc)

	if words != 15 {
		t.Fatalf("Wrong word count. Expected: %d, Got: %d", 12, words)
	}
	if images != 1 {
		t.Fatalf("Wrong image count. Expected %d, Got: %d", 1, words)
	}
}
