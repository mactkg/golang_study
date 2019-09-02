package main

import (
	"bytes"
	"testing"
)

func TestGraph(t *testing.T) {
	original := bytes.Buffer{}
	new := bytes.Buffer{}

	graphOriginal(&original)
	graphNew(&new)

	if original.String() != new.String() {
		t.Fatal()
	}
}
