package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"
)

// color from https://color.adobe.com/ja/My-Color-Theme-color-theme-12907198
var palette = []color.Color{
	color.RGBA{0x80, 0x68, 0x53, 0xFF},
	color.RGBA{0xFF, 0xCF, 0xA6, 0xFF},
	color.RGBA{0xCC, 0xA6, 0x85, 0xFF},
}

const (
	baseIndex   = 0
	lineIndex   = 1
	accentIndex = 2
)

func lissajous(cycles float64, out io.Writer) {
	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)

			// 何らかの興味深い方法
			if cycles*math.Pi*0.5 < t && t < cycles*math.Pi*1.5 {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), lineIndex)
			} else {
				img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), accentIndex)
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

// server
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var cycles = 5.0
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		if c := r.FormValue("cycles"); c != "" {
			i, _ := strconv.Atoi(c)
			cycles = float64(i)
		}
		lissajous(cycles, w)
	})

	log.Println("Try difference!: http://localhost:8000\thttp://localhost:8000?cycles=2\thttp://localhost:8000?cycles=10")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
