package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	width, height = 600, 320
	cells         = 50
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	file, err := os.Create("./eggBox.svg")
	if err != nil {
		fmt.Println("file open error", err)
	}
	writeSVG(file, eggBox)
	file.Close()

	file, err = os.Create("./hump.svg")
	if err != nil {
		fmt.Println("file open error", err)
	}
	writeSVG(file, hump)
	file.Close()

	file, err = os.Create("./moth.svg")
	if err != nil {
		fmt.Println("file open error", err)
	}
	writeSVG(file, moth)
	file.Close()
}

type fn func(float64, float64) float64

func writeSVG(out io.Writer, f fn) {
	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			fmt.Fprintf(out, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, f fn) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	// https://gyazo.com/f8b5cf337bffabcf70eaa40a9bc9bb3a.png
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func hump(x, y float64) float64 {
	a := math.Sin(y)
	b := math.Sin(x)
	return math.Pow(2, a) * math.Pow(2, b) / 12
}

func moth(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(-x) * math.Pow(1.5, -r)
}

func eggBox(x, y float64) float64 {
	return math.Sin(x*y/10) / 10
}
