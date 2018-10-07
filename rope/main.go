package main

//go:generate go get github.com/usedbytes/hsv
// https://github.com/fogleman/gg

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"github.com/usedbytes/hsv"
)

type modul struct {
	modul int
	add   int
	start int
}

type seq struct {
	snd  bool
	x, y int
}

type handler func(s1, s2 seq)

func main() {
	log.Println("This is", os.Args[0], "(C) 2018 by Martin Lercher")

	html := strings.Replace(root, "#'", "`", -1)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(html))
	})

	http.HandleFunc("/rope.png", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), 500)
			return
		}
		wi := atoi(r, "w", 31)
		add := atoi(r, "a", 1)
		m := modul{wi, add, 1}
		img := m.render()
		png.Encode(w, img)
	})

	go func(url string) {
		log.Println("starting browser at", url)
		openBrowser(url)
	}(fmt.Sprintf("http://localhost:%d/", 8000))

	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func (m modul) render() *image.NRGBA {
	rect := image.Rect(0, 0, m.modul, m.modul)
	img := image.NewNRGBA(rect)
	i := 0

	log.Println(m)
	m.iterate(func(s1, s2 seq) {
		//log.Println(s1, s2)
		if s1.x == s2.x && s1.y == s2.y {
			i++
			c := toHSV(i)
			set(img, s1.x, s1.y, c)
		} else if s1.x == s2.x {
			step := sgn(s1.y, s2.y)
			for y := s1.y; y != s2.y; y += step {
				i++
				c := toHSV(i)
				set(img, s1.x, y, c)
			}
		} else {
			step := sgn(s1.x, s2.x)
			for x := s1.x; x != s2.x; x += step {
				i++
				c := toHSV(i)
				set(img, x, s1.y, c)
			}
		}
	})
	return img
}

func set(img *image.NRGBA, x, y int, c color.Color) {
	img.Set(x-1, y-1, c)
	img.Set(x, y-1, c)
	img.Set(x+1, y-1, c)
	img.Set(x-1, y, c)
	img.Set(x, y, c)
	img.Set(x+1, y, c)
	img.Set(x-1, y+1, c)
	img.Set(x, y+1, c)
	img.Set(x+1, y+1, c)
}

func sgn(i, j int) int {
	if i <= j {
		return 1
	}
	return -1
}

func toHSV(i int) color.Color {
	return hsv.HSVColor{H: uint16((i>>3) % 360), S: 255, V: 255}
}

func (m modul) iterate(h handler) {
	c0 := m.findCycle()
	c1 := m.step(c0)
	// log.Println("cycle start", c0, c1)
	s0 := seq{true, c1, c0}
	s1, s2 := s0, s0
	for {
		c1 = m.step(c1)
		s1, s2 = s2, s2.append(c1)
		h(s1, s2)
		if s0.eq(s2) {
			return
		}
	}
}

func (m modul) findCycle() int {
	var sq, ss int = m.start, m.start
	for {
		ss = m.step(ss)
		sq = m.step(m.step(sq))
		if sq == ss {
			return ss
		}
	}
}

func (m modul) step(state int) int {
	s1 := state*state + m.add
	return s1 % m.modul
}

func (s seq) append(i int) seq {
	if s.snd {
		return seq{false, s.x, i}
	}
	return seq{true, i, s.y}
}

func (s seq) eq(s1 seq) bool {
	return s.x == s1.x && s.y == s1.y
}

func (s seq) String() string {
	return fmt.Sprintf("[%4d, %4d]", s.x, s.y)
}

// see https://gist.github.com/threeaccents/607f3bc3a57a2ddd9d57
// openBrowser tries to open the URL in a browser,
// and returns whether it succeed in doing so.
func openBrowser(url string) bool {
	var args []string
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open", url}
	case "windows":
		args = []string{"cmd", "/c", "start", strings.Replace(url, "&", "^&", -1)}
	default:
		args = []string{"xdg-open", url}
	}
	cmd := exec.Command(args[0], args[1:]...)
	return cmd.Start() == nil
}

func atoi(r *http.Request, p string, d int) int {
	it := r.FormValue(p)
	i, err := strconv.Atoi(it)
	if err != nil {
		return d
	}
	return i
}

const root = `<html>
<head>
	<title>Ropes</title>
</head>
<body style="background-color: #444; color: #ddd;">
	<h1>
		Colorful ropes
	</h1>
	<p>
	  <img id="img">
	  <br/>
	  <a href="#" id="info" target="_blank" style="color: #7ff;"></a>
	</p>
	<script>
const size = 801;
let a = 0;
const img = document.getElementById('img');
const info = document.getElementById('info');

function ld() {
	a++;
	a = a % size;
	const src = #'rope.png?w=${size}&a=${a}#';
	info.innerText = #'f(x) = x^2 + ${a} mod ${size}#';
	info.href = src;
	img.width = size;
	img.height = size;
	img.src = src;
	console.log('loading:', src);
	setTimeout(ld, 500);
}

ld();

	</script>
</body>
</html>`
