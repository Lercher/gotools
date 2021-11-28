# ell

ell plots "near-zero solutions" of the elliptic equation

y^2 = x^3 - x + 0.45

On my hardware, evaluating the functional value in 64bit
float takes under 2ns, so colorful high-res images should be
possible to make in short succession.

```txt
goos: windows
goarch: amd64
pkg: github.com/lercher/gotools/ell
cpu: Intel(R) Core(TM) i7-8700 CPU @ 3.20GHz
BenchmarkElliptic-12    	702969988	         1.705 ns/op	       0 B/op	       0 allocs/op
```