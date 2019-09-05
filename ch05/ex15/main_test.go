package main

import (
	"testing"
)

func TestMax(t *testing.T) {
	res := max(3, 2, 1, 3, 2)
	if res != 3 {
		t.Fatalf("Expected: %v, Got: %v", 3, res)
	}

	res = max(-3, -2, 3, 1)
	if res != 3 {
		t.Fatalf("Expected: %v, Got: %v", 3, res)
	}

	res = max()
	if res != 0 {
		t.Fatalf("Expected: %v, Got: %v", 0, res)
	}
}

func TestMin(t *testing.T) {
	res := min(3, 2, 1, 3, 2)
	if res != 1 {
		t.Fatalf("Expected: %v, Got: %v", 1, res)
	}

	res = min(-3, -2, 3, 1)
	if res != -3 {
		t.Fatalf("Expected: %v, Got: %v", -3, res)
	}

	res = min()
	if res != 0 {
		t.Fatalf("Expected: %v, Got: %v", 0, res)
	}
}

func TestMax2(t *testing.T) {
	res := max2(-100)
	if res != -100 {
		t.Fatalf("Expected: %v, Got: %v", -100, res)
	}

	res = max2(10, 5, 20)
	if res != 20 {
		t.Fatalf("Expected: %v, Got: %v", 20, res)
	}
}

func TestMin2(t *testing.T) {
	res := min2(100)
	if res != 100 {
		t.Fatalf("Expected: %v, Got: %v", 100, res)
	}

	res = min2(10, 5, 20)
	if res != 5 {
		t.Fatalf("Expected: %v, Got: %v", 5, res)
	}
}
