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

var res int

func BenchmarkPopCountSingle(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcount.PopCountSingle(10 * uint64(i))
	}
	res = temp
}

func BenchmarkPopCountLoop(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcount.PopCountLoop(10 * uint64(i))
	}
	res = temp
}

func BenchmarkPopCountLoopShift(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcount.PopCountLoopShift(10 * uint64(i))
	}
	res = temp
}

func BenchmarkPopCountLoopWithOutShift(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += popcount.PopCountLoopWithOutShift(10 * uint64(i))
	}
	res = temp
}
