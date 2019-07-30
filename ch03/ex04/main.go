package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"strconv"
)

type renderOption struct {
	width    int
	height   int
	cells    int
	xyrange  float64
	radAngle float64
}

func (opt *renderOption) xyscale() float64 {
	return float64(opt.width) / 2 / opt.xyrange
}

func (opt *renderOption) zscale() float64 {
	return float64(opt.height) * 0.4
}

func (opt *renderOption) calcRadAngles() (float64, float64) {
	return math.Sin(opt.radAngle), math.Cos(opt.radAngle)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		option := &renderOption{
			width:    600,
			height:   320,
			cells:    100,
			xyrange:  30.0,
			radAngle: math.Pi / 6.0,
		}
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		if w := r.FormValue("width"); w != "" {
			i, _ := strconv.Atoi(w)
			option.width = i
		}
		if h := r.FormValue("height"); h != "" {
			i, _ := strconv.Atoi(h)
			option.height = i
		}
		if c := r.FormValue("cells"); c != "" {
			i, _ := strconv.Atoi(c)
			option.cells = i
		}
		if r := r.FormValue("xyrange"); r != "" {
			f, _ := strconv.ParseFloat(r, 64)
			option.xyrange = f
		}
		if r := r.FormValue("radAngle"); r != "" {
			f, _ := strconv.ParseFloat(r, 64)
			option.radAngle = f
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		render(*option, w)
	})
	log.Println("Try differences!\n" +
		"http://localhost:8000/?width=800&height=800&cells=50&radAngle=0.3925\n" +
		"http://localhost:8000/?width=200&height=200&cells=20&radAngle=0.3925\n")
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func render(opt renderOption, out io.Writer) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width: 0.7;' "+
		"width='%d' height='%d'>", opt.width, opt.height)

	// calculate maximum / minimum height for coloring
	max, min := math.SmallestNonzeroFloat64, math.MaxFloat64
	for i := 0; i < opt.cells; i++ {
		for j := 0; j < opt.cells; j++ {
			x := opt.xyrange * (float64(i)/float64(opt.cells) - 0.5)
			y := opt.xyrange * (float64(j)/float64(opt.cells) - 0.5)
			v := f(float64(x), float64(y))
			max = math.Max(max, v)
			min = math.Min(min, v)
		}
	}
	length := max - min

	for i := 0; i < opt.cells; i++ {
		for j := 0; j < opt.cells; j++ {
			// position
			ax, ay := corner(i+1, j, opt)
			bx, by := corner(i, j, opt)
			cx, cy := corner(i, j+1, opt)
			dx, dy := corner(i+1, j+1, opt)

			// coloring
			x := opt.xyrange * (float64(i)/float64(opt.cells) - 0.5)
			y := opt.xyrange * (float64(j)/float64(opt.cells) - 0.5)
			absz := (f(float64(x), float64(y)) - min) / length
			var r, g, b float64
			if absz > 0.5 {
				r = 255 * (absz - 0.5) * 2
			} else {
				b = 255 * absz * 2
			}

			// generate polygon
			fmt.Fprintf(out, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' "+
				"style='stroke: rgb(%f, %f, %f);' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, opt renderOption) (float64, float64) {
	sin, cos := opt.calcRadAngles()

	x := opt.xyrange * (float64(i)/float64(opt.cells) - 0.5)
	y := opt.xyrange * (float64(j)/float64(opt.cells) - 0.5)
	z := f(x, y)

	// https://gyazo.com/f8b5cf337bffabcf70eaa40a9bc9bb3a.png
	sx := float64(opt.width)/2 + (x-y)*cos*opt.xyscale()
	sy := float64(opt.height)/2 + (x+y)*sin*opt.xyscale() - z*opt.zscale()

	return sx, sy
}

func f(x, y float64) float64 {
	if x == 0 && y == 0 {
		return 1
	}

	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
