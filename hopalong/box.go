package main

type box struct {
	x, y line
}

func (b *box) enclose(x, y float64) {
	b.x.enclose(x)
	b.y.enclose(y)
}

func (b *box) transform(w, h int, quadrant1 bool) *transform {
	t := &transform{
		x: b.x.linetransform(w, quadrant1),
		y: b.y.linetransform(h, quadrant1),
	}
	t.keepAspectRatio()
	return t
}