package main

import (
	"testing"
)

func BenchmarkBadEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BadEcho()
	}
}

func BenchmarkGoodEcho(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GoodEcho()
	}
}
