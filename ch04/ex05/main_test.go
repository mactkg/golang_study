package main

import "testing"

func TestRemoveDupNeighbor(t *testing.T) {
	a := []string{"ab", "abc", "abc", "def", "def", "def", "abc"}
	expected := []string{"ab", "abc", "def", "abc"}
	a = RemoveDupNeighbor(a)

	for i := range a {
		if a[i] != expected[i] {
			t.Fatalf("Error! Got: %s, Expected: %s", a[i], expected[i])
		}
	}
}

func TestRemoveDupNeighborMinimumCase(t *testing.T) {
	a := []string{"ab", "ab"}
	expected := []string{"ab"}
	a = RemoveDupNeighbor(a)

	for i := range a {
		if a[i] != expected[i] {
			t.Fatalf("Error! Got: %s, Expected: %s", a[i], expected[i])
		}
	}
}
