package main

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestElementsByTagName(t *testing.T) {
	f, err := os.Open("./test.html")
	if err != nil {
		t.Fatal(err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		t.Fatal(err)
	}

	res := ElementByTagName(doc, "body")
	if len(res) != 1 {
		t.Fatalf("<body> count | expected: 1, got: %v", len(res))
	}

	res = ElementByTagName(doc, "undefined")
	if len(res) != 0 {
		t.Fatalf("<undefined> shouldn't be found, but got: %v", len(res))
	}

	res = ElementByTagName(doc, "img", "head", "h1", "h2")
	for _, v := range res {
		switch v.Data {
		case "img", "head", "h1", "h2":
			// nop
		default:
			t.Fatalf("shound't be leached")
		}
	}

	res = ElementByTagName(doc)
	if len(res) != 0 {
		t.Fatal()
	}
}
