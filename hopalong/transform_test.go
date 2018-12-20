package main

import "testing"

func TestTransform(t *testing.T) {
	b := &box{}
	b.enclose(-4, -1)
	b.enclose(4, 3)
	t.Log(*b)
	tr := b.transform(800, 600, false) // all quadrants
	t.Log(*tr.x, *tr.y)
	want(t, tr, -4, -1, 0, 0)
	want(t, tr, 3.99, 2.99, 799, 399)
	want(t, tr, 0, 0, 400, 100)
}

func TestTransformQ1only(t *testing.T) {
	b := &box{}
	b.enclose(-4, -1)
	b.enclose(4, 3)
	t.Log(*b)
	tr := b.transform(800, 600, true) // only quadrant 1 (top left)
	t.Log(*tr.x, *tr.y)
	wantnok(t, tr, -4, -1, -800, -200)
	want(t, tr, 3.99, 0, 798, 0)
	want(t, tr, 0, 2.99, 0, 598)
	want(t, tr, 0, 0, 0, 0)
}

func want(t *testing.T, tr *transform, x, y float64, wa, wb int) {
	t.Helper()
	a, b, ok := tr.tr(x, y)
	if !ok {
		t.Errorf("tr(%v,%v) want (%v, %v), got (%v, %v) and not ok", x, y, wa, wb, a, b)
	}
	if ok && (a != wa || b != wb) {
		t.Errorf("tr(%v,%v) want (%v, %v), got (%v, %v) with ok", x, y, wa, wb, a, b)
	}
}

func wantnok(t *testing.T, tr *transform, x, y float64, wa, wb int) {
	t.Helper()
	a, b, ok := tr.tr(x, y)
	if ok {
		t.Errorf("tr(%v,%v) want (%v, %v), got (%v, %v) and ok", x, y, wa, wb, a, b)
	}
	if !ok && (a != wa || b != wb) {
		t.Errorf("tr(%v,%v) want (%v, %v), got (%v, %v) with ok", x, y, wa, wb, a, b)
	}
}
