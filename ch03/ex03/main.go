package main

import (
	"fmt"
	"math"
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
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width: 0.7;' "+
		"width='%d' height='%d'>", width, height)

	// calculate maximum / minimum height for coloring
	max, min := math.SmallestNonzeroFloat64, math.MaxFloat64
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			v := f(float64(x), float64(y))
			max = math.Max(max, v)
			min = math.Min(min, v)
		}
	}
	length := max - min

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// position
			ax, ay := corner(i+1, j)
			bx, by := corner(i, j)
			cx, cy := corner(i, j+1)
			dx, dy := corner(i+1, j+1)

			// coloring
			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			absz := (f(float64(x), float64(y)) - min) / length
			var r, g, b float64
			if absz > 0.5 {
				r = 255 * (absz - 0.5) * 2
			} else {
				b = 255 * absz * 2
			}

			// generate polygon
			fmt.Printf("<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' "+
				"style='stroke: rgb(%f, %f, %f);' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	// https://gyazo.com/f8b5cf337bffabcf70eaa40a9bc9bb3a.png
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale

	return sx, sy
}

func f(x, y float64) float64 {
	if x == 0 && y == 0 {
		return 1
	}

	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
