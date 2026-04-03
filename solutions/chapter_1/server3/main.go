// http responses handler

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand/v2"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/lissajous", lissajous)
	log.Fatal(http.ListenAndServe("localhost:8889", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}

var alpha uint8 = 0

const (
	backgroundIndex = 0 // first color of palette
	blackIndex      = 1 // next color of palette
)

func lissajous(out http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	cycles := 10 // count of full oscillations
	if countStr := r.FormValue("count"); countStr != "" {
		if parsed, err := strconv.Atoi(countStr); err == nil {
			cycles = parsed
		}
	}

	const (
		res         = 0.001 // angle resolutuion
		size        = 100   // size of image
		nframes int = 256   // frames count
		delay       = 8     // delay between frames (one unit is 10ms)
	)
	freq := rand.Float64() * 3.0 // frequency of oscillation
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for frame := range nframes {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		getGetNextAlpha(frame, nframes)
		var palette = []color.Color{color.RGBA{alpha, alpha, 0x00, alpha}, color.Black}
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < float64(cycles)*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func getGetNextAlpha(frame int, nframes int) {
	if frame < int(nframes/3) {
		alpha += uint8(256 / nframes)
	} else {
		alpha -= uint8(256 / nframes)
	}
}
