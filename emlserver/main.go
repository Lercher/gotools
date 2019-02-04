package main

// go build && emlserver.exe -d \\ntsdtsg\c$\inetpub\mailroot\Drop
// go run *.go -d drop

import (
	"flag"
	"fmt"
	"log"
	"mime"
	"net/http"
	"os"
)

var (
	flagDropDir = flag.String("d", `\inetpub\mailroot\Drop`, "directory containing *.eml files")
	flagPort    = flag.Int("port", 9090, "port number of the webserver")
	flagRoot    = flag.String("root", "email", "root path for the eml web server")
)

var worddecoder = new(mime.WordDecoder)

func main() {
	log.Println("This is emlServer, (C) 2019 by Martin Lercher")
	flag.Parse()

	stat, err := os.Stat(*flagDropDir)
	if err != nil {
		log.Fatalln(err)
	}
	if !stat.IsDir() {
		log.Fatalln(*flagDropDir, "must be a directory containing 0..n *.eml files")
	}
	srv := &server{*flagRoot, *flagDropDir}
	ht := fmt.Sprintf(":%d", *flagPort)
	log.Println("serving *.eml files from", srv.directory, "on", ht, srv.prefix())

	http.Handle(srv.prefix(), srv)
	log.Fatalln(http.ListenAndServe(ht, srv))
}
