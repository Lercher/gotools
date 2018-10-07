package main

//go:generate go get github.com/usedbytes/hsv

import (
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/png"
	"log"
	"math"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/usedbytes/hsv"
)

const root = `<html>
<head>
	<script src="https://cdn.rawgit.com/anvaka/panzoom/v4.4.0/dist/panzoom.min.js"></script>
</head>
<body>
	<h1>
		Mandelbrot
		<select id="sel">
			<option value="pal">Plan9 palette</option>
			<option value="hsv1">HSV 1</option>
			<option value="hsv2">HSV 2</option>
			<option value="hsv3">HSV 3</option>
			<option value="hsv4">HSV 4</option>
		</select>
	</h1>
	<p>
      <img id="img">
	</p>
	<p>
  		<a id="a2" href="#" target="_blank">Image twice as big and with more depth</a>
		<br/>
		<a id="a4" href="#" target="_blank">Image for times as big and with more depth</a>
	</p>
	<script>
const zoom1 = 400;

let loc = {
	x: -0.67,
	y: 0,
	z: 1,
	d: 200,
	w: 800,
	h: 600,
	i: 0
}
let lod;

const img = document.getElementById('img');
const a2 = document.getElementById('a2');
const a4 = document.getElementById('a4');
const sel = document.getElementById('sel');
let pz;

document.body.addEventListener('zoom',   handler, true);
document.body.addEventListener('panend', handler, true);
sel.addEventListener('change', handler, true);

function handler(e) {
	const tr = pz.getTransform();
	panloc(tr.x, tr.y, tr.scale);

	loc.i++;
	const i = loc.i;
	setTimeout(function() {ld(i);}, 250);
}


function panloc(x, y, sc) {
	console.log(x, y, sc);
	loc.z = lod.z / sc;
	const sc1 = (sc - 1);
	dx = (x + x + img.width  * sc1) * lod.z / zoom1;
	dy = (y + y + img.height * sc1) * lod.z / zoom1;
	loc.x = lod.x - dx;
	loc.y = lod.y - dy;
}

img.addEventListener('load', function(e) {		
	pz = panzoom(img);
});

function ld(i) {
	if (i === loc.i) {
		if (pz) 
			pz.dispose();
		lod = JSON.parse(JSON.stringify(loc));
		let mt = sel.value;
		const src = #'mb.png?m=${mt}&w=${loc.w}&h=${loc.h}&x=${loc.x}&y=${loc.y}&z=${loc.z}&d=${loc.d}#';
		img.width = loc.w;
		img.height = loc.h;
		img.src = src;
		a2.href = #'mb.png?m=${mt}&w=${loc.w * 2}&h=${loc.h * 2}&x=${loc.x}&y=${loc.y}&z=${loc.z / 2}&d=${loc.d * 10}#';
		a4.href = #'mb.png?m=${mt}&w=${loc.w * 4}&h=${loc.h * 4}&x=${loc.x}&y=${loc.y}&z=${loc.z / 4}&d=${loc.d * 10}#';
		console.log('loading:', src);
	}
}

ld(loc.i);
	</script>
</body>
</html>`

var pal = palette.Plan9
var nt = runtime.NumCPU()
var cm = map[string]mapper{
	"pal":  toRGBA,
	"hsv1": toHSV1,
	"hsv2": toHSV2,
	"hsv3": toHSV3,
	"hsv4": toHSV4,
}

// http://localhost:8000/

func main() {
	log.Println("This is", os.Args[0], "(C) 2018 by Martin Lercher")
	ch := make(chan line)
	renderServer(ch)

	html := strings.Replace(root, "#'", "`", -1)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	})

	http.HandleFunc("/mb.png", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		wi := atoi(r, "w", 32)
		he := atoi(r, "h", 32)
		de := atoi(r, "d", 100)
		center := complex(atof(r, "x", -0.5), atof(r, "y", 0))
		zoom := atof(r, "z", 16) / 400
		method := r.FormValue("m")
		plus := complex(float64(wi)*zoom, float64(he)*zoom)
		z1 := center - plus
		z2 := center + plus
		t := time.Now()
		img := renderPaletted(ch, wi, he, de, z1, z2, method)
		log.Println(wi, he, de, z1, z2, time.Now().Sub(t))
		w.Header().Add("Content-Type", "image/png")
		png.Encode(w, img)
	})

	log.Println("Starting browser on port ", 8000, "...")
	go func(url string) {
		openBrowser(url)
	}(fmt.Sprintf("http://localhost:%d", 8000))

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

type line struct {
	wg     *sync.WaitGroup
	img    *image.NRGBA
	y      int
	w      int
	depth  int
	zline  complex128
	dx     complex128
	mapper mapper
}

func renderServer(ch <-chan line) {
	for i := 0; i < nt; i++ {
		go func() {
			for l := range ch {
				l.render()
			}
		}()
	}
	log.Println("set up", nt, "render workers")
}

func renderPaletted(ch chan<- line, w, h, depth int, z1, z2 complex128, method string) *image.NRGBA {
	rect := image.Rect(0, 0, w, h)
	img := image.NewNRGBA(rect)

	mapper, ok := cm[method]
	if !ok {
		mapper = toRGBA
	}

	dz := z2 - z1
	dx := complex(real(dz)/float64(w), 0)
	dy := complex(0, imag(dz)/float64(h))
	var wg sync.WaitGroup
	wg.Add(h)
	hh := h / nt
	for n := 0; n < nt; n++ {
		from, to := n*hh, n*hh+hh
		if n+1 == nt {
			to = h
		}
		for y := from; y < to; y++ {
			zline := z1 + scalarmult(y, dy)
			l := line{&wg, img, y, w, depth, zline, dx, mapper}
			ch <- l
		}
	}
	wg.Wait()
	return img
}

func (l line) render() {
	for x := 0; x < l.w; x++ {
		z := l.zline + scalarmult(x, l.dx)
		idx := l.mapper(iter(z, l.depth))
		l.img.Set(x, l.y, idx)
	}
	l.wg.Done()
}

func (l line) rendert() {
	t := time.Now()
	l.render()
	log.Println(l.y, time.Now().Sub(t))
}

func scalarmult(n int, z complex128) complex128 {
	s := float64(n)
	return complex(s*real(z), s*imag(z))
}

type mapper func(iter, maxiter int, abs2 float64) color.Color

func toHSV1(iter, maxiter int, abs2 float64) color.Color {
	return toHSV(iter, maxiter, 0, abs2)
}

func toHSV2(iter, maxiter int, abs2 float64) color.Color {
	return toHSV(iter, maxiter, 90, abs2)
}

func toHSV3(iter, maxiter int, abs2 float64) color.Color {
	return toHSV(iter, maxiter, 180, abs2)
}

func toHSV4(iter, maxiter int, abs2 float64) color.Color {
	return toHSV(iter, maxiter, 270, abs2)
}

func toHSV(iter, maxiter, ofs int, abs2 float64) color.Color {
	if iter >= maxiter {
		return color.Black
	}
	smooth := float64(iter+ofs) - math.Log(math.Log(abs2))
	smoothint := uint16(int(smooth) % 360)
	return hsv.HSVColor{H: smoothint, S: 255, V: 255}
}

func toRGBA(iter, maxiter int, abs2 float64) color.Color {
	if iter >= maxiter {
		return color.Black
	}
	smooth := float64(iter) - math.Log(math.Log(abs2))
	smoothindex := int(smooth)
	smoothremainder := smooth - float64(smoothindex)
	if smoothindex < 0 {
		smoothindex = 0
	}
	if smoothremainder < 0 {
		smoothremainder = 0
	}
	idx1 := smoothindex % (len(pal) - 1)
	idx1++
	idx2 := idx1 + 1
	if idx2 >= len(pal) {
		idx2 = 1
	}
	c1, c2 := pal[idx1], pal[idx2]
	return interpolate(c1, c2, smoothremainder)
}

func interpolate(c1, c2 color.Color, f float64) color.Color {
	r1, g1, b1, _ := c1.RGBA()
	r2, g2, b2, _ := c2.RGBA()
	return color.RGBA{m(r1, r2, f), m(g1, g2, f), m(b1, b2, f), 255}
}

func m(a, b uint32, f float64) uint8 {
	// a/b 0..ffff, f 0..1
	p := f*float64(b) + (1-f)*float64(a)
	p /= 256
	return uint8(p)
}
func iter(c complex128, maxiter int) (int, int, float64) {
	var z complex128
	var a float64
	for i := 0; i < maxiter; i++ {
		z = z*z + c
		a = abs2(z)
		if a > 4 {
			return i, maxiter, a
		}
	}
	return maxiter, maxiter, a
}

func abs2(z complex128) float64 {
	r, i := real(z), imag(z)
	return r*r + i*i
}

func atoi(r *http.Request, p string, d int) int {
	it := r.FormValue(p)
	i, err := strconv.Atoi(it)
	if err != nil {
		return d
	}
	return i
}

func atof(r *http.Request, p string, d float64) float64 {
	it := r.FormValue(p)
	i, err := strconv.ParseFloat(it, 64)
	if err != nil {
		return d
	}
	return i
}

// see https://gist.github.com/threeaccents/607f3bc3a57a2ddd9d57
// openBrowser tries to open the URL in a browser,
// and returns whether it succeed in doing so.
func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
	default:
		args = []string{"xdg-open"}
	}
	cmd := exec.Command(args[0], append(args[1:], url)...)
	return cmd.Start() == nil
}
