package main

import (
	"image"
	"image/color"
	"image/png"
	"io"
)

func hopPNG(w io.Writer, width, height, rounds int) {
	img := hopimg(width, height, rounds)
	png.Encode(w, img)
}

func hopimg(width, height, rounds int) *image.NRGBA {
	h := &hop{}
	h.randomizeABC()
	h.bounds(rounds)
	t := h.transform(width, height)
	h.reset()

	c := color.RGBA{255, 30, 10, 255}
	rect := image.Rect(0, 0, width, height)
	img := image.NewNRGBA(rect)

	h.rounds(rounds, func() {
		x, y := t.tr(h.x, h.y)
		img.Set(x, y, c)
	})

	return img
}
