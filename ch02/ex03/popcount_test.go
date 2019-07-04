package main

import (
	"testing"

	"github.com/mactkg/golang_study/ch02/ex03/popcount"
)

func TestPopCounts(t *testing.T) {
	single := popcount.PopCountSingle(10)
	loop := popcount.PopCountLoop(10)
	if single != loop {
		t.Fatalf("failed: %v != %v", single, loop)
	}
}

func BenchmarkPopCountSingle(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountSingle(10)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoop(10)
	}
}
