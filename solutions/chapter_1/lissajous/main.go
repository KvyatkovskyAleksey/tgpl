// lissajous generate animated GIF from random lissajous figures
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

var alpha uint8 = 0

const (
	backgroundIndex = 0 // first color of palette
	blackIndex      = 1 // next color of palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles      = 10    // count of full oscillations
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
		for t := 0.0; t < cycles*2*math.Pi; t += res {
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
