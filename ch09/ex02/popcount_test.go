package popcount

import (
	"testing"
)

func TestPopCounts(t *testing.T) {
	single := PopCountSingle(10)
	loopShift := PopCountLoopWithOutShift(10)
	if single != loopShift {
		t.Fatalf("failed: %v != %v", single, loopShift)
	}
}

var res int

func BenchmarkPopCountSingle(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += PopCountSingle(10 * uint64(i))
	}
	res = temp
}

func BenchmarkPopCountSingleLazy(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += PopCountSingleLazy(10 * uint64(i))
	}
	res = temp
}

func BenchmarkPopCountLoop(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += PopCountLoop(10 * uint64(i))
	}
	res = temp
}

func BenchmarkPopCountLoopShift(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += PopCountLoopShift(10 * uint64(i))
	}
	res = temp
}

func BenchmarkPopCountLoopWithOutShift(b *testing.B) {
	temp := 0
	for i := 0; i < b.N; i++ {
		temp += PopCountLoopWithOutShift(10 * uint64(i))
	}
	res = temp
}
