package main

// go get && go run main.go x-info hello  x-to world  <main.go >sample.txt
// go get && go run main.go -r < sample.txt

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/textproto"
	"os"
)

var (
	flagRead = flag.Bool("r", false, "read mode, i.e. validate a file and log it's header information")
)

func main() {
	log.Println("This is", os.Args[0], "- a program to read and write http protocol files")
	flag.Parse()

	if *flagRead {
		if flag.NArg() != 0 {
			log.Println("usage:", os.Args[0], "-r")
			log.Println("  reads from stdin, checks and writes header information to stdout")
			flag.Usage()
			os.Exit(2)
		}
		read()
		os.Exit(0)
	}

	if flag.NArg() == 0 || flag.NArg()%2 != 0 {
		log.Println("usage:", os.Args[0], "(<key> <value>)+")
		log.Println("  adds the headers specified and then copies stdin to the body")
		flag.Usage()
		os.Exit(2)
	}
	write()
}

func read() {
	buf := bufio.NewReader(os.Stdin)
	rd := textproto.NewReader(buf)

	h, err := rd.ReadMIMEHeader()
	if err != nil {
		log.Fatalln(err, "while reading header")
	}
	hdr := http.Header(h)
	hdr.Write(os.Stdout)

	n, err := io.Copy(ioutil.Discard, buf)
	if err != nil {
		log.Fatalln(err, "while reading body")
	}
	log.Println(n, "bytes body")
}

func write() {
	hdr := make(http.Header)
	for i := 0; i < flag.NArg(); i += 2 {
		key := flag.Args()[i]
		val := flag.Args()[i+1]
		hdr.Add(key, val)
	}

	err := hdr.Write(os.Stdout)
	if err != nil {
		log.Fatalln(err, "while writing header")
	}
	fmt.Fprintln(os.Stdout)
	n, err := io.Copy(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatalln(err, "while writing body")
	}
	log.Println(n, "bytes body written")
}
