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

	supersample := make([][]color.Color, width*2)
	for i := 0; i < width*2; i++ {
		supersample[i] = make([]color.Color, height*2)
	}
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for py := 0; py < height*2; py++ {
		y := float64(py)/(height*2)*(ymax-ymin) + ymin // ymax ~ ymin
		for px := 0; px < width*2; px++ {
			x := float64(px)/(width*2)*(xmax-xmin) + xmin // xmax ~ xmin
			z := complex(x, y)
			supersample[px][py] = mandelbrot(z)
		}
	}

	for py := 0; py < height; py++ {
		for px := 0; px < width; px++ {
			ar, ag, ab, _ := supersample[2*px][2*py].RGBA()
			br, bg, bb, _ := supersample[2*px][2*py+1].RGBA()
			cr, cg, cb, _ := supersample[2*px+1][2*py].RGBA()
			dr, dg, db, _ := supersample[2*px+1][2*py+1].RGBA()

			img.Set(px, py, color.RGBA{
				uint8((ar + br + cr + dr) / 4),
				uint8((ag + bg + cg + dg) / 4),
				uint8((ab + bb + cb + db) / 4),
				0xFF,
			})
		}
	}

	png.Encode(os.Stdout, img)
}

func mandelbrot(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 { // out of "circle"
			return color.Gray{255 - contrast*n}
		}
	}
	return color.Black // you are in mandelbrot!
}
