package main

import (
	"strings"
	"testing"
)

func TestJoin(t *testing.T) {
	expected := strings.Join([]string{"A", "B", "C"}, "-")
	got := join("-", "A", "B", "C")
	if expected != got {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}

	expected = strings.Join([]string{}, "-")
	got = join("-")
	if expected != got {
		t.Fatalf("expected: %v, got: %v", expected, got)
	}
}
