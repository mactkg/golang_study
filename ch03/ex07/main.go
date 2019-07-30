package main

import (
	"image"
	"image/color"
	"image/png"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 1024, 1024
	)

	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin // ymax ~ ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin // xmax ~ xmin
			z := complex(x, y)
			img.Set(px, py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

/*
 *  x_(n+1) = x_n - (f(x_n) / f'(x_n))
 *          => z - (z^4 - 1) / 4z^3
 *			= z - (z - 1/z^3) / 4
 */
func mandelbrot(z complex128) color.Color {
	const iterations = 100
	const contrast = 255 / 60

	var v = z
	for n := uint8(0); n < iterations; n++ {
		v = v - (v-1/(v*v*v))/4
		if cmplx.Abs(v*v*v*v-1) < 1e-5 { // out of "circle"
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black // you are in mandelbrot!
}
