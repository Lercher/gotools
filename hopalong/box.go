package main

type box struct {
	x, y line
}

func (b *box) enclose(x, y float64) {
	b.x.enclose(x)
	b.y.enclose(y)
}