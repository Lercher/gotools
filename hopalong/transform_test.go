package main

import "testing"

func TestTransform(t *testing.T) {
	b := &box{}
	b.enclose(-4, -1)
	b.enclose(4, 3)
	t.Log(*b)
	tr := b.transform(800, 600)
	t.Log(*tr.x, *tr.y)
	want(t, tr, -4, -1, 0, 0)
	want(t, tr, 4, 3, 800, 400)
	want(t, tr, 0, 0, 400, 100)
}

func want(t *testing.T, tr *transform, x, y float64, wa, wb int) {
	t.Helper()
	a, b := tr.tr(x, y)
	if a != wa || b != wb {
		t.Errorf("tr(%v,%v) want (%v, %v), got (%v, %v)", x, y, wa, wb, a, b)
	}
}
