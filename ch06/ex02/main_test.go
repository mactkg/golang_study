package main

import "testing"

func TestAddAll(t *testing.T) {
	var x IntSet

	x.AddAll(10, 100, 42)
	x.AddAll(80, 80, 80, 80)

	if x.Len() != 4 {
		t.Fatalf("Expected: %d, Got: %d", 4, x.Len())
	}
}
