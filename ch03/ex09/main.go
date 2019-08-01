package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/cmplx"
	"net/http"
	"strconv"
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
}

func (opt *renderOption) calcX(px float64) float64 {
	return px/opt.width*(opt.xmax-opt.xmin) + opt.xmin // xmax ~ xmin
}

func (opt *renderOption) calcY(py float64) float64 {
	return py/opt.height*(opt.ymax-opt.ymin) + opt.ymin
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		option := renderOption{
			width:      512,
			height:     512,
			xmin:       -2,
			ymin:       -2,
			xmax:       +2,
			ymax:       +2,
			iterations: 50,
			contrast:   20,
		}
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}

		var x, y, s = 0., 0., 1.
		if v := r.FormValue("x"); v != "" {
			i, _ := strconv.ParseFloat(v, 64)
			x = i
		}
		if v := r.FormValue("y"); v != "" {
			i, _ := strconv.ParseFloat(v, 64)
			y = i
		}
		if v := r.FormValue("scale"); v != "" {
			f, _ := strconv.ParseFloat(v, 64)
			s = f
		}
		x = x / (option.width / 4)
		y = y / (option.height / 4)
		s = 2.0 / s
		option.xmin = -s + x
		option.xmax = s + x
		option.ymin = -s + y
		option.ymax = s + y

		img := image.NewRGBA(image.Rect(0, 0, int(option.width), int(option.height)))
		DrawMangelbrot128(img, option)
		png.Encode(w, img)
	})
	fmt.Println("Try it out:\n",
		"http://localhost:8000/?scale=1&y=80&x=46\n",
		"http://localhost:8000/?scale=5000&y=80&x=46\n",
		"http://localhost:8000/?scale=10000&y=80&x=46",
	)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// DrawMangelbrot128 draws mandelbrot image using complex128
func DrawMangelbrot128(img *image.RGBA, opt renderOption) {
	fmt.Printf("DrawMangelbrot128: %v\n", opt)
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
	}
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
