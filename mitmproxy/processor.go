package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
)

type processor struct {
	c   chan rr
	fld string
}

type rr struct {
	Request  *http.Request
	Response *http.Response
	Body     []byte
}

func (p *processor) fm() map[string]any {
	return map[string]any{
		"match": path.Match,
		"splitfn": func(p string) string {
			_, fn := path.Split(p)
			return fn
		},
		"splitdir": func(p string) string {
			dir, _ := path.Split(p)
			return dir
		},
		"join":    path.Join,
		"folder":  p.folder,
		"foldern": p.foldern,
		"save":    p.save,
	}
}

func (p *processor) foldern(pth string, lastN int) (string, error) {
	f := foldernPath(pth, lastN)
	return p.folder(f)
}

func foldernPath(pth string, lastN int) string {
	pth = strings.TrimSuffix(pth, "/")
	ps := strings.Split(pth, "/")
	if lastN > len(ps) || lastN <= 0 {
		log.Println("ERROR: foldern(", pth, lastN, "): has", len(ps), "parts, wants last", lastN)
	}
	ps = ps[len(ps)-lastN:]
	return path.Join(ps...)
}

func (p *processor) folder(f string) (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	f = path.Join(filepath.ToSlash(wd), f)
	s := fmt.Sprintf("using folder: %v", f)
	p.fld = f
	err = os.MkdirAll(p.fld, 0644)
	return s, err
}

func (p *processor) send(r *http.Request, resp *http.Response) {
	if resp.StatusCode != http.StatusOK {
		// log.Println("NOT queued response:", resp.Status)
		return
	}

	bs, err := io.ReadAll(resp.Body)
	resp.Body = io.NopCloser(bytes.NewBuffer(bs))
	if err != nil {
		log.Println("reading response body of", r.URL, ":", err)
		return
	}

	p.c <- rr{Request: r, Response: resp, Body: bs}
}

func (p *processor) save(fn string, body []byte) string {
	pth := path.Join(p.fld, fn)
	pth = filepath.FromSlash(pth)

	err := os.WriteFile(pth, body, 0644)
	if err != nil {
		log.Println("NOT saved response to", pth, "write:", err)
		return ""
	}

	return fmt.Sprintf("saved %v byte(s) of response to %v", len(body), pth)
}

func (p *processor) start(tplFile string) error {
	t, err := template.New("").Funcs(p.fm()).ParseFiles(tplFile)
	if err != nil {
		return err
	}

	p.c = make(chan rr, 1)
	go func() {
		err := t.ExecuteTemplate(log.Default().Writer(), "main", p.c)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	return nil
}
