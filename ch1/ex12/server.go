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

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	keys := r.URL.Query()
	cycles, err := strconv.Atoi(keys["cycles"][0])
	if err != nil {
		log.Print(err)
	}
	lissajous(w, float64(cycles))
}

func lissajous(out io.Writer, cycles float64) {
	var palette = []color.Color{color.RGBA{0x00, 0x00, 0x00, 0xff}, color.RGBA{0x00, 0xff, 0x00, 0xff}, color.RGBA{0xff, 0x00, 0x00, 0xff}}
	const (
		blackIndex = 0 // パレットの最初の色
		greenIndex = 1 // パレットの次の色
		redIndex   = 2 // パレットの3色目
	)
	const (
		res     = 0.001 // 回転の分解能
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
			var colorIndex uint8
			if t < cycles*2*math.Pi/2 {
				colorIndex = greenIndex
			} else {
				colorIndex = redIndex
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
