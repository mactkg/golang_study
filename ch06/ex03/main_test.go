package main

import "testing"

func TestIntersectWith(t *testing.T) {
	var s, n IntSet

	s.AddAll(10, 15, 80)
	n.AddAll(10, 30)

	s.IntersectWith(&n)
	if s.Len() != 1 || !s.Has(10) {
		t.Fatalf("Expected: %v, Got: %v", 1, s.Len())
	}
}

func TestDifferenceWith(t *testing.T) {
	var s, n IntSet

	s.AddAll(10, 15, 80)
	n.AddAll(10, 30)

	s.DifferenceWith(&n)
	if s.Len() != 2 || !s.Has(15) || !s.Has(80) {
		t.Fatalf("Expected: %v, Got: %v", 2, s.Len())
	}
}

func TestSymmetricDifferenceWith(t *testing.T) {
	var s, n IntSet

	s.AddAll(10, 15, 80)
	n.AddAll(10, 30)

	s.SymmetricDifference(&n)
	if s.Len() != 3 && !s.Has(15) && !s.Has(30) && !s.Has(80) {
		t.Fatalf("Expected: %v, Got: %v", 3, s.Len())
	}
}
