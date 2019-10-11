package main

import (
	"bytes"
	"io"
	"testing"
)

/*

goos: darwin
goarch: amd64
pkg: github.com/mactkg/golang_study/ch08/ex05
Benchmark1-8                  52          20726167 ns/op
Benchmark2-8                  70          18903531 ns/op
Benchmark4-8                  67          17473057 ns/op
Benchmark8-8                  61          17399514 ns/op * <- sweet spot
Benchmark16-8                 61          17329071 ns/op
Benchmark32-8                 58          21224696 ns/op
PASS
ok      github.com/mactkg/golang_study/ch08/ex05        8.737s

 */

var buf io.Writer
func Benchmark1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf = &bytes.Buffer{}
		writeSVG(buf, 1)
	}
}

func Benchmark2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf = &bytes.Buffer{}
		writeSVG(buf, 2)
	}
}

func Benchmark4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf = &bytes.Buffer{}
		writeSVG(buf, 4)
	}
}

func Benchmark8(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf = &bytes.Buffer{}
		writeSVG(buf, 8)
	}
}

func Benchmark16(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf = &bytes.Buffer{}
		writeSVG(buf, 16)
	}
}

func Benchmark32(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf = &bytes.Buffer{}
		writeSVG(buf, 1)
	}
}
