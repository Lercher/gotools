package main

import (
	"testing"
)

func Test_processor_fm1(t *testing.T) {
	p := &processor{}
	m := p.fm()["match"].(func(pattern, name string) (matched bool, err error))
	ok, err := m("/webtoon/*/*/", "/webtoon/life-0003/chapter-2/")
	if err != nil {
		t.Error("err: ", err)
	}
	if !ok {
		t.Error("no match")
	}
}

func Test_processor_fm2(t *testing.T) {
	p := &processor{}
	m := p.fm()["match"].(func(pattern, name string) (matched bool, err error))
	ok, err := m("/chapters/*/*/*.jpg", "/chapters/manga_5d2c7626253fbf429/b947937c16feab5c0595844e9/02.jpg")
	if err != nil {
		t.Error("err: ", err)
	}
	if !ok {
		t.Error("no match")
	}
}

func Test_processor_foldernPath(t *testing.T) {
	have := "/webtoon/life-0003/chapter-2/"
	want := "life-0003/chapter-2"	
	got := foldernPath(have, 2)
	if got != want {
		t.Errorf("foldernPath(%v, 2): want %v, got %v", have, want, got)
	}
}

func Test_processor_foldernPath2(t *testing.T) {
	have := "/webtoon/life-0003/chapter-2"
	want := "life-0003/chapter-2"	
	got := foldernPath(have, 2)
	if got != want {
		t.Errorf("foldernPath(%v, 2): want %v, got %v", have, want, got)
	}
}
