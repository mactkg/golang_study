package main

import (
	"fmt"
	"io"
	"math"
	"os"
	"sync"
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
	writeSVG(os.Stdout, 1)
}

func writeSVG(w io.Writer, parallel int) {
	p := make(chan struct{}, parallel)

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='fill: white; stroke-width: 0.7;' "+
		"width='%d' height='%d'>", width, height)

	// calculate maximum / minimum height for coloring
	max, min := math.SmallestNonzeroFloat64, math.MaxFloat64
	wg := sync.WaitGroup{}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			wg.Add(1)
			p <- struct{}{}
			go func(i, j int) {
				x := xyrange * (float64(i)/cells - 0.5)
				y := xyrange * (float64(j)/cells - 0.5)
				v := f(float64(x), float64(y))
				max = math.Max(max, v)
				min = math.Min(min, v)
				<-p
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	length := max - min

	mutex := sync.Mutex{}
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			wg.Add(1)
			p <- struct{}{}
			go func(i, j int) {
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

				mutex.Lock()
				// generate polygon
				fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g' "+
				"style='stroke: rgb(%f, %f, %f);' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, r, g, b)
				mutex.Unlock()
				<-p
				wg.Done()
			}(i, j)
		}
	}
	wg.Wait()
	fmt.Fprintln(w, "</svg>")
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
