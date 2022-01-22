package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
)

type forward struct {
	u *url.URL
}

func (fwd forward) proxyall(w http.ResponseWriter, r *http.Request) {
	c := http.Client{}
	u, err := url.Parse(r.URL.String()) // copy URL
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	u.Host = fwd.u.Host
	u.Scheme = fwd.u.Scheme

	log.Println(r.URL, "->", u)
	resp, err := c.Get(u.String())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	w.WriteHeader(resp.StatusCode)
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		log.Println(r.URL, "content", err.Error())
		return
	}
}

var flagRE = flag.String("save", "*.jpg", "save all url.paths matching this glob pattern to the output dir")
var flagOut = flag.String("o", "./output", "directory to put files into. It will be created if it won't exist")

func main() {
	log.Println("This is imgsave, (C) 2022 by Martin Lercher")
	log.Println("It proxies GET queries to the http/s://host:port argument and saves all responses")
	log.Println("matching a glob pattern (-save pattern) as files during browsing localhost")

	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		fmt.Println("\nhttp/s://host argument missing")
		os.Exit(1)
	}

	u, err := url.Parse(flag.Arg(0))
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("Relaying to", u)

	fwd := forward{
		u: u,
	}

	addr := ":9000"
	log.Println("Listening on http://localhost" + addr)

	http.HandleFunc("/", fwd.proxyall)
	http.ListenAndServe(addr, nil)
}
