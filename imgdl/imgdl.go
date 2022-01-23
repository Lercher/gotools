package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"text/template"
	"time"

	"golang.design/x/clipboard"
)

var (
	flagOut  = flag.String("o", "./output", "directory to put files into. It will be created if it won't exist")
	flagStop = flag.Bool("stoponerror", true, "stop loop on first http get error")
)

var (
	clipver  = 0
	lastclip = string(clipboard.Read(clipboard.FmtText))
)

func main() {
	log.Println("This is imgdl, (C) 2022 by Martin Lercher")
	log.Println("It http GETs web ressources provided by a sequence of urls")
	log.Println("URLs are provided by a Go text template called 'main' line by line")
	log.Println("and responses are saved flattened in the outpupt directory (-o)")
	log.Println(``)
	log.Println(`{{define "main"}}{{range $i := intRange 1 12}}file_{{$i}}.png{{end}}{{end}} might be handy here`)
	log.Println(`{{$star := star "abc***def**ghi" "123abc789de234"}} -> ["abc" "de"] by position`)
	log.Println(`{{$cliboardtext := clip}} -> read clipboard`)
	log.Println(`{{waitclip}} -> busy wait for clipboard text to change`)
	log.Println(``)

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		fmt.Println("\ntemplate container file name missing")
		os.Exit(1)
	}

	t, err := template.New("").Funcs(funcs).ParseFiles(flag.Args()[0])
	if err != nil {
		log.Fatalln(err)
	}

	dir := *flagOut
	err = os.MkdirAll(dir, 0777)
	if err != nil {
		log.Fatalln(err)
	}

	pr, pw := io.Pipe()
	go func() {
		err := t.ExecuteTemplate(pw, "main", nil)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	scanner := bufio.NewScanner(pr)
	currentclipver := clipver
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		if clipver == currentclipver {
			log.Println("SKIPING", clipver, line)
			continue
		}
		if u, err := url.Parse(line); err == nil {
			cont := func() bool {
				_, fn := path.Split(u.Path)
				fullname := path.Join(dir, fn)
				fullname = filepath.FromSlash(fullname)
				resp, err := http.Get(u.String())
				if err != nil {
					log.Println(u, err)
					return false
				}
				defer resp.Body.Close()
				if resp.StatusCode != http.StatusOK {
					log.Println(u, resp.Status)
					return false
				}

				f, err := os.Create(fullname)
				if err != nil {
					log.Println(u, err)
					return false
				}
				defer f.Close()

				n, err := io.Copy(f, resp.Body)
				if err != nil {
					log.Println(u, resp.Status)
					return false
				}

				err = f.Close()
				if err != nil {
					log.Println(u, resp.Status)
					return false
				}

				fmt.Println(n, "bytes written to", fullname, "|", line)
				return true
			}()
			if !cont {
				currentclipver = clipver
			}
		} else {
			log.Println("IGNORED:", err)
		}
	}
}

var funcs = template.FuncMap{
	"intRange": func(start, end int) []int {
		n := end - start + 1
		result := make([]int, n)
		for i := 0; i < n; i++ {
			result[i] = start + i
		}
		return result
	},
	"star": func(startemplate, s string) ([]string, error) {
		startemplate = strings.TrimSpace(startemplate)
		s = strings.TrimSpace(s)
		if len(startemplate) > len(s) {
			return nil, fmt.Errorf("template containing '*' is longer than the string to be split: %v > %v", len(startemplate), len(s))
		}
		var sa []string
		sr := []rune(s)
		var sb strings.Builder
		for i, ch := range startemplate {
			if ch == '*' {
				sb.WriteRune(sr[i])
			} else {
				if sb.Len() > 0 {
					sa = append(sa, sb.String())
				}
				sb = strings.Builder{}
			}
		}
		log.Println("star yields", sa)
		return sa, nil
	},
	"clip": func() string {
		return string(clipboard.Read(clipboard.FmtText))
	},
	"waitclip": func() string {		
		fmt.Print("Waiting for clipboard to change...")
		for {
			<-time.After(250 * time.Millisecond)
			clp := string(clipboard.Read(clipboard.FmtText))
			if lastclip != clp {
				lastclip = clp
				break
			}
			fmt.Print(".")
		}
		fmt.Println(" OK")
		clipver++
		return ""
	},
}
