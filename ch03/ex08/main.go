package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"math"
	"math/big"
	"math/cmplx"
	"os"
)

type renderOption struct {
	width      float64
	height     float64
	xmin       float64
	ymin       float64
	xmax       float64
	ymax       float64
	xOffset    float64
	yOffset    float64
	iterations uint8
	contrast   uint8
	debug      bool
}

func (opt *renderOption) calcX(px float64) float64 {
	return px/opt.width*(opt.xmax-opt.xmin) + opt.xmin // xmax ~ xmin
}

func (opt *renderOption) calcY(py float64) float64 {
	return py/opt.height*(opt.ymax-opt.ymin) + opt.ymin
}

func main() {
	option := renderOption{
		width:      32,
		height:     32,
		xmin:       -2,
		ymin:       -2,
		xmax:       +2,
		ymax:       +2,
		iterations: 10,
		contrast:   20,
		debug:      true,
	}
	img := image.NewRGBA(image.Rect(0, 0, int(option.width*2), int(option.height*2)))

	option.xOffset = option.width
	option.yOffset = option.height
	DrawMangelbrotBigRat(img, option)

	option.xOffset = 0
	option.yOffset = 0
	DrawMangelbrot64(img, option)

	option.xOffset = option.width
	option.yOffset = 0
	DrawMangelbrot128(img, option)

	option.xOffset = 0
	option.yOffset = option.height
	DrawMangelbrotBigFloat(img, option)

	file, err := os.Create("./output.png")
	if err != nil {
		fmt.Println("file open error", err)
	}
	defer file.Close()
	png.Encode(file, img)
}

// DrawMangelbrot64 draws mandelbrot image using complex64
func DrawMangelbrot64(img *image.RGBA, opt renderOption) {
	wait := make(chan bool)

	for py := 0.; py < opt.height; py++ {
		y := opt.calcY(py)
		for px := 0.; px < opt.width; px++ {
			x := opt.calcX(px)

			go func(px, py int, x, y float64) {
				z64 := complex(float32(x), float32(y))
				img.Set(px, py, mandelbrot64(z64, opt.iterations, opt.contrast))
				wait <- true
			}(int(px+opt.xOffset), int(py+opt.yOffset), x, y)
		}
	}

	for i := 0.; i < opt.width*opt.height; i++ {
		<-wait
		if opt.debug && int(i)%int(math.Floor(opt.width*opt.height/10)) == 0 {
			fmt.Printf("64: %d(%.3f)\n", int(i), i/(opt.width*opt.height))
		}
	}
}

// DrawMangelbrot128 draws mandelbrot image using complex128
func DrawMangelbrot128(img *image.RGBA, opt renderOption) {
	wait := make(chan bool)

	for py := 0.; py < opt.height; py++ {
		y := opt.calcY(py)
		for px := 0.; px < opt.width; px++ {
			x := opt.calcX(px)

			go func(px, py int, x, y float64) {
				z64 := complex(x, y)
				img.Set(px, py, mandelbrot128(z64, opt.iterations, opt.contrast))
				wait <- true
			}(int(px+opt.xOffset), int(py+opt.yOffset), x, y)
		}
	}

	for i := 0.; i < opt.width*opt.height; i++ {
		<-wait
		if opt.debug && int(i)%int(math.Floor(opt.width*opt.height/10)) == 0 {
			fmt.Printf("128: %d(%.3f)\n", int(i), i/(opt.width*opt.height))
		}
	}
}

// DrawMangelbrotBigFloat draws mandelbrot image using big.Float
func DrawMangelbrotBigFloat(img *image.RGBA, opt renderOption) {
	wait := make(chan bool)

	for py := 0.; py < opt.height; py++ {
		y := opt.calcY(py)
		for px := 0.; px < opt.width; px++ {
			x := opt.calcX(px)

			go func(px, py int, x, y float64) {
				bigFloatX, bigFloatY := big.NewFloat(x), big.NewFloat(y)

				img.Set(px, py, mandelbrotBigFloat(bigFloatX, bigFloatY, opt.iterations, opt.contrast))
				wait <- true
			}(int(px+opt.xOffset), int(py+opt.yOffset), x, y)
		}
	}

	for i := 0.; i < opt.width*opt.height; i++ {
		<-wait
		if opt.debug && int(i)%int(math.Floor(opt.width*opt.height/10)) == 0 {
			fmt.Printf("big.Float: %d(%.3f)\n", int(i), i/(opt.width*opt.height))
		}
	}
}

// DrawMangelbrotBigRat draws mandelbrot image using big.Rat
func DrawMangelbrotBigRat(img *image.RGBA, opt renderOption) {
	wait := make(chan bool)

	for py := 0.; py < opt.height; py++ {
		y := opt.calcY(py)
		for px := 0.; px < opt.width; px++ {
			x := opt.calcX(px)

			go func(px, py int, x, y float64) {
				bigRatX, bigRatY := big.NewRat(0, 1), big.NewRat(0, 1)
				bigRatX.SetFloat64(x)
				bigRatY.SetFloat64(y)

				img.Set(px, py, mandelbrotBigRat(bigRatX, bigRatY, opt.iterations, opt.contrast))
				wait <- true
			}(int(px+opt.xOffset), int(py+opt.yOffset), x, y)
		}
	}

	for i := 0.; i < opt.width*opt.height; i++ {
		<-wait
		if opt.debug && int(i)%int(math.Floor(opt.width*opt.height/10)) == 0 {
			fmt.Printf("big.Rat: %d(%.3f)\n", int(i), i/(opt.width*opt.height))
		}
	}
}

func mandelbrot64(z complex64, iterations, contrast uint8) color.Color {
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

func mandelbrot128(z complex128, iterations, contrast uint8) color.Color {
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

func mandelbrotBigFloat(zReal, zImag *big.Float, iterations, contrast uint8) color.Color {
	l := big.NewFloat(2)

	vReal, vImag := new(big.Float), new(big.Float)
	for n := uint8(0); n < iterations; n++ {
		// v = v*v + z
		vReal, vImag = SquareBigFloatComplex(vReal, vImag)
		vReal.Add(vReal, zReal)
		vImag.Add(vImag, zImag)

		if AbsBigFloatComplex(vReal, vImag).Cmp(l) == 1 { // abs > 2, out of "circle"
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

func mandelbrotBigRat(zReal, zImag *big.Rat, iterations, contrast uint8) color.Color {
	l := big.NewRat(2, 1)

	vReal, vImag := new(big.Rat), new(big.Rat)
	for n := uint8(0); n < iterations; n++ {
		// v = v*v + z
		vReal, vImag = SquareBigRatComplex(vReal, vImag)
		vReal.Add(vReal, zReal)
		vImag.Add(vImag, zImag)

		if AbsBigRatComplex(vReal, vImag).Cmp(l) == 1 { // abs > 2, out of "circle"
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
