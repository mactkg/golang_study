package main

import (
	"bytes"
	"testing"

	"golang.org/x/net/html"
)

func TestPrettyPrint(t *testing.T) {
	result := bytes.Buffer{}
	err := prettyPrint("http://gopl.io", &result)
	if err != nil {
		t.Fatal()
	}

	_, err = html.Parse(&result)
	if err != nil {
		t.Fatal(err)
	}
}
