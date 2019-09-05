package main

import (
	"os"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettyPrint(t *testing.T) {
	f, err := os.Open("./test.html")
	if err != nil {
		t.Fatal(err)
	}

	doc, err := html.Parse(f)
	if err != nil {
		t.Fatal(err)
	}
	res := ElementByID(doc, "toc")
	if res == nil {
		t.Fatal("#toc should be found")
	}

	res = ElementByID(doc, "Foo")
	if res != nil {
		t.Fatalf("#Foo shoudn't be found. but got %v", res)
	}
}
