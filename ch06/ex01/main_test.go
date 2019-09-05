package main

import "testing"

func TestLen(t *testing.T) {
	var x IntSet

	if x.Len() != 1 {
		t.Fatalf("Expected: %d, Got: %d", 0, x.Len())
	}

	x.Add(10)
	x.Add(100)
	x.Add(42)

	if x.Len() != 3 {
		t.Fatalf("Expected: %d, Got: %d", 3, x.Len())
	}
}

func TestRemove(t *testing.T) {
	var x IntSet

	x.Add(10)
	x.Add(23)
	x.Add(34)

	x.Remove(10)

	if x.Len() != 2 {
		t.Fatalf("Expected: %d, Got: %d", 2, x.Len())
	}
}

func TestClear(t *testing.T) {
	var x IntSet

	x.Add(10)
	x.Add(20)
	x.Add(130)

	x.Clear()

	if x.Len() != 0 {
		t.Fatalf("Expected: %d, Got: %d", 0, x.Len())
	}
}

func TestCopy(t *testing.T) {
	var x IntSet

	x.Add(10)
	x.Add(100)
	x.Add(180)

	y := x.Copy()

	y.Clear()

	if x.Len() != 3 {
		t.Fatalf("Expected: %d, Got: %d", 3, x.Len())
	}

	if y.Len() != 0 {
		t.Fatalf("Expected: %d, Got: %d", 0, y.Len())
	}
}
