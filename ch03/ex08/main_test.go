package main

import (
	"image"
	"testing"
)

var option = renderOption{
	width:      16,
	height:     16,
	xmin:       -2,
	ymin:       -2,
	xmax:       +2,
	ymax:       +2,
	iterations: 20,
	contrast:   20,
}
var img = image.NewRGBA(image.Rect(0, 0, int(option.width*2), int(option.height*2)))

func BenchmarkDrawMangelbrot64(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DrawMangelbrot64(img, option)
	}
}

func BenchmarkDrawMangelbrot128(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DrawMangelbrot128(img, option)
	}
}

func BenchmarkDrawMangelbrotBigFloat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DrawMangelbrotBigFloat(img, option)
	}
}

func BenchmarkDrawMangelbrotBigRat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		DrawMangelbrotBigRat(img, option)
	}
}
