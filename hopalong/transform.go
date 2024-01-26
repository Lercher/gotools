package main

import "math"

type transform struct {
	x, y *linetransform
}

func (t *transform) keepAspectRatio() {
	f := math.Min(t.x.factor, t.y.factor)
	t.x.factor = f
	t.y.factor = f
}

func (t *transform) tr(x, y float64) (int, int, bool) {
	ix, okx := t.x.tr(x)
	iy, oky := t.y.tr(y)
	return ix, iy, okx && oky
}
