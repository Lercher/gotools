package main

type box struct {
	x, y line
}

func (b *box) enclose(x, y float64) {
	b.x.enclose(x)
	b.y.enclose(y)
}

func (b *box) transform(w, h int) *transform {
	t := &transform{
		x: b.x.linetransform(w),
		y: b.y.linetransform(h),
	}
	t.keepAspectRatio()
	return t
}