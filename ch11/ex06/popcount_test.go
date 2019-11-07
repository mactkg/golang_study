package main

import (
	"m/popcount"
	"testing"
)

var res int

func benchmarkPopCountSingle(b *testing.B, size int) {
	for j := 0; j < b.N; j++ {
		temp := 0
		for i := 0; i < size; i++ {
			temp += popcount.PopCountSingle(10 * uint64(i))
		}
		res = temp
	}
}
func BenchmarkPopCountSingle10(b *testing.B)       { benchmarkPopCountSingle(b, 10) }
func BenchmarkPopCountSingle100(b *testing.B)      { benchmarkPopCountSingle(b, 100) }
func BenchmarkPopCountSingle1000(b *testing.B)     { benchmarkPopCountSingle(b, 1000) }
func BenchmarkPopCountSingle10000(b *testing.B)    { benchmarkPopCountSingle(b, 10000) }
func BenchmarkPopCountSingle100000(b *testing.B)   { benchmarkPopCountSingle(b, 100000) }
func BenchmarkPopCountSingle1000000(b *testing.B)  { benchmarkPopCountSingle(b, 1000000) }
func BenchmarkPopCountSingle10000000(b *testing.B) { benchmarkPopCountSingle(b, 10000000) }

func benchmarkPopCountLoop(b *testing.B, size int) {
	for j := 0; j < b.N; j++ {
		temp := 0
		for i := 0; i < size; i++ {
			temp += popcount.PopCountLoop(10 * uint64(i))
		}
		res = temp
	}
}
func BenchmarkPopCountLoop10(b *testing.B)       { benchmarkPopCountLoop(b, 10) }
func BenchmarkPopCountLoop100(b *testing.B)      { benchmarkPopCountLoop(b, 100) }
func BenchmarkPopCountLoop1000(b *testing.B)     { benchmarkPopCountLoop(b, 1000) }
func BenchmarkPopCountLoop10000(b *testing.B)    { benchmarkPopCountLoop(b, 10000) }
func BenchmarkPopCountLoop100000(b *testing.B)   { benchmarkPopCountLoop(b, 100000) }
func BenchmarkPopCountLoop1000000(b *testing.B)  { benchmarkPopCountLoop(b, 1000000) }
func BenchmarkPopCountLoop10000000(b *testing.B) { benchmarkPopCountLoop(b, 10000000) }

func benchmarkPopCountLoopShift(b *testing.B, size int) {
	for j := 0; j < b.N; j++ {
		temp := 0
		for i := 0; i < size; i++ {
			temp += popcount.PopCountLoopShift(10 * uint64(i))
		}
		res = temp
	}
}
func BenchmarkPopCountLoopShift10(b *testing.B)       { benchmarkPopCountLoopShift(b, 10) }
func BenchmarkPopCountLoopShift100(b *testing.B)      { benchmarkPopCountLoopShift(b, 100) }
func BenchmarkPopCountLoopShift1000(b *testing.B)     { benchmarkPopCountLoopShift(b, 1000) }
func BenchmarkPopCountLoopShift10000(b *testing.B)    { benchmarkPopCountLoopShift(b, 10000) }
func BenchmarkPopCountLoopShift100000(b *testing.B)   { benchmarkPopCountLoopShift(b, 100000) }
func BenchmarkPopCountLoopShift1000000(b *testing.B)  { benchmarkPopCountLoopShift(b, 1000000) }
func BenchmarkPopCountLoopShift10000000(b *testing.B) { benchmarkPopCountLoopShift(b, 10000000) }

func benchmarkPopCountLoopWithOutShift(b *testing.B, size int) {
	for j := 0; j < b.N; j++ {
		temp := 0
		for i := 0; i < size; i++ {
			temp += popcount.PopCountLoopWithOutShift(10 * uint64(i))
		}
		res = temp
	}
}
func BenchmarkPopCountLoopWithOutShift10(b *testing.B)    { benchmarkPopCountLoopWithOutShift(b, 10) }
func BenchmarkPopCountLoopWithOutShift100(b *testing.B)   { benchmarkPopCountLoopWithOutShift(b, 100) }
func BenchmarkPopCountLoopWithOutShift1000(b *testing.B)  { benchmarkPopCountLoopWithOutShift(b, 1000) }
func BenchmarkPopCountLoopWithOutShift10000(b *testing.B) { benchmarkPopCountLoopWithOutShift(b, 10000) }
func BenchmarkPopCountLoopWithOutShift100000(b *testing.B) {
	benchmarkPopCountLoopWithOutShift(b, 100000)
}
func BenchmarkPopCountLoopWithOutShift1000000(b *testing.B) {
	benchmarkPopCountLoopWithOutShift(b, 1000000)
}
func BenchmarkPopCountLoopWithOutShift10000000(b *testing.B) {
	benchmarkPopCountLoopWithOutShift(b, 10000000)
}
