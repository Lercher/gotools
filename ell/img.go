package main

import (
	"image"
	"image/png"
	"io"
	"math"

	"github.com/usedbytes/hsv"
)

func ellPNG(wr io.Writer, w float64, steps int, epsilon float64) error {
	rect := image.Rect(0, 0, steps, steps)
	img := image.NewNRGBA(rect)
	imgwalk(img, w, steps, epsilon)
	return png.Encode(wr, img)
}

func imgwalk(img *image.NRGBA, w float64, steps int, epsilon float64) {
	czo := hsv.HSVColor{H: 15, S: 150, V: 40} // H: uint16(0..359)
	walk00(w, steps, epsilon, func(i, j int, _, _, v float64, isZero bool) {
		if isZero {
			img.Set(i, j, czo)
		} else {
			// lg := math.Log(math.Abs(v))
			lg360 := math.Trunc(math.Remainder(v*160, 360))
			cnz := hsv.HSVColor{H: uint16(lg360), S: 180, V: 255} // H: uint16(0..359)
			img.Set(i, j, cnz)
		}
	})
}
