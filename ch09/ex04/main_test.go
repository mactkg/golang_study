package main

import (
	"testing"
)

func Benchmark10(b *testing.B) {
	in, out := challenge(10)
	for i := 0; i < b.N; i++ {
		in<-i
		<-out
	}
}

func Benchmark100(b *testing.B) {
	in, out := challenge(100)
	for i := 0; i < b.N; i++ {
		in<-i
		<-out
	}
}

func Benchmark1000(b *testing.B) {
	in, out := challenge(1000)
	for i := 0; i < b.N; i++ {
		in<-i
		<-out
	}
}

func Benchmark10000(b *testing.B) {
	in, out := challenge(10000)
	for i := 0; i < b.N; i++ {
		in<-i
		<-out
	}
}

func Benchmark100000(b *testing.B) {
	in, out := challenge(100000)
	for i := 0; i < b.N; i++ {
		in<-i
		<-out
	}
}

func Benchmark1000000(b *testing.B) {
	in, out := challenge(1000000)
	for i := 0; i < b.N; i++ {
		in<-i
		<-out
	}
}

