package main

import (
	"image"
	"image/color"
	"image/png"
	"io"

	"github.com/usedbytes/hsv"
)

func hopPNG(w io.Writer, width, height, rounds, nextColor int, quadrant1 bool) {
	img := hopimg(width, height, rounds, nextColor, quadrant1)
	png.Encode(w, img)
}

func hopimg(width, height, rounds, nextColor int, quadrant1 bool) *image.NRGBA {
	h := &hop{}
	h.randomizeABC()
	h.bounds(rounds)
	t := h.transform(width, height, quadrant1)
	h.reset()

	rect := image.Rect(0, 0, width, height)
	img := image.NewNRGBA(rect)

	n := nextColor + 1
	ci := 0
	var c color.Color
	h.rounds(rounds, func() {
		n++
		if n >= nextColor {
			n = 0
			ci++
			c = hsv.HSVColor{H: uint16(ci % 360), S: 255, V: 255}
		}
		x, y, ok := t.tr(h.x, h.y)
		if ok {
			img.Set(x, height-y-1, c)
		}
	})

	return img
}
