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

func (t *transform) tr(x, y float64) (int, int) {
	return t.x.tr(x), t.y.tr(y)
}
