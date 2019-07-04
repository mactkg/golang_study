package main

import (
	"testing"

	"github.com/mactkg/golang_study/ch02/ex05/popcount"
)

func TestPopCounts(t *testing.T) {
	single := popcount.PopCountSingle(10)
	loopShift := popcount.PopCountLoopWithOutShift(10)
	if single != loopShift {
		t.Fatalf("failed: %v != %v", single, loopShift)
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

func BenchmarkPopCountLoopShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoopShift(10)
	}
}

func BenchmarkPopCountLoopWithOutShift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountLoopWithOutShift(10)
	}
}
