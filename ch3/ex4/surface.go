package main

import (
	"fmt"
	"io"
	"math"
	"net/http"
	"strconv"
)

const (
	cells   = 100
	xyrange = 30.0

	angle = math.Pi / 6

	minZ = -1.
	maxZ = 1.
)

var width, height, xyscale, zscale float64
var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var err error

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		width, err = strconv.ParseFloat(r.FormValue("width"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		height, err = strconv.ParseFloat(r.FormValue("height"), 64)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "image/svg+xml")
		surface(w)
	}
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8000", nil)
}

func surface(w io.Writer) {
	xyscale = width / 2 / xyrange
	zscale = height * 0.4
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)

			red, blue := setColor(az, bz, cz, dz)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='rgb(%d,0,%d)' />\n",
				ax, ay, bx, by, cx, cy, dx, dy, red, blue)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func setColor(a, b, c, d float64) (int, int) {
	m := (a + b + c + d) / 4
	red := int(255 * (m - minZ) / (maxZ - minZ))
	blue := 255 - red
	return red, blue
}
