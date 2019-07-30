package main

import (
	"image"
	"image/color"
	"image/png"
	"math/big"
	"math/cmplx"
	"os"
)

func main() {
	const (
		xmin, ymin, xmax, ymax = -2, -2, +2, +2
		width, height          = 512, 512
		allWidth, allHeight    = width * 2, height * 2
	)

	img := image.NewRGBA(image.Rect(0, 0, allWidth, allHeight))
	for py := 0; py < height; py++ {
		y := float64(py)/height*(ymax-ymin) + ymin // ymax ~ ymin
		for px := 0; px < width; px++ {
			x := float64(px)/width*(xmax-xmin) + xmin // xmax ~ xmin

			z64 := complex(float32(x), float32(y))
			z128 := complex(x, y)
			img.Set(px, py, mandelbrot64(z64))
			img.Set(width+px, py, mandelbrot128(z128))
			// img.Set(px, height+py, mandelbrot(z))
			// img.Set(width+px, height+py, mandelbrot(z))
		}
	}
	png.Encode(os.Stdout, img)
}

func mandelbrot64(z complex64) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex64
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(complex128(v)) > 2 { // out of "circle"
			switch n % 3 {
			case 0:
				return color.RGBA{255 - contrast*n, 128, 128, 0xff}
			case 1:
				return color.RGBA{128, 255 - contrast*n, 128, 0xff}
			case 2:
				return color.RGBA{128, 128, 255 - contrast*n, 0xff}
			}
		}
	}
	return color.Black // you are in mandelbrot!
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 200
	const contrast = 15

	var v complex128
	for n := uint8(0); n < iterations; n++ {
		v = v*v + z
		if cmplx.Abs(v) > 2 { // out of "circle"
			switch n % 3 {
			case 0:
				return color.RGBA{255 - contrast*n, 128, 128, 0xff}
			case 1:
				return color.RGBA{128, 255 - contrast*n, 128, 0xff}
			case 2:
				return color.RGBA{128, 128, 255 - contrast*n, 0xff}
			}
		}
	}
	return color.Black // you are in mandelbrot!
}

// xx = (a + ib)(a + ib) = (aa - bb) + i(ab + ba)
func mandelbrotBigFloat(zReal, zImag *big.Float) color.Color {
	const iterations = 200
	const contrast = 15

	vReal, vImag := new(big.Float), new(big.Float)
	for n := uint8(0); n < iterations; n++ {
		vReal, vImag = multiplyBigFloatComplex(vReal, vImag, vReal, vImag)
		vReal.Add(vReal, zReal)
		vImag.Add(vImag, zImag)

		p := new(big.Float)
		p.Add(new(big.Float).Mul(vReal, vReal), new(big.Float).Mul(vReal, vReal))

		v, _ := p.Sqrt(p).Float64()
		if v > 2 { // out of "circle"
			switch n % 3 {
			case 0:
				return color.RGBA{255 - contrast*n, 128, 128, 0xff}
			case 1:
				return color.RGBA{128, 255 - contrast*n, 128, 0xff}
			case 2:
				return color.RGBA{128, 128, 255 - contrast*n, 0xff}
			}
		}
	}
	return color.Black // you are in mandelbrot!
}

func multiplyBigFloatComplex(aReal, aImag, bReal, bImag *big.Float) (*big.Float, *big.Float) {
	var r, i, a, b big.Float
	return r.Sub(a.Mul(aReal, bReal), b.Mul(aImag, bImag)), i.Add(a.Mul(aReal, bImag), b.Mul(aImag, bReal))
}
