package main

import (
	"image"
	"image/color"
	"image/png"
	"io"

	"github.com/usedbytes/hsv"
)

func hopPNG(w io.Writer, width, height, rounds, nextColor int) {
	img := hopimg(width, height, rounds, nextColor)
	png.Encode(w, img)
}

func hopimg(width, height, rounds, nextColor int) *image.NRGBA {
	h := &hop{}
	h.randomizeABC()
	h.bounds(rounds)
	t := h.transform(width, height)
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
		x, y := t.tr(h.x, h.y)
		img.Set(x, y, c)
	})

	return img
}
