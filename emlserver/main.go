package main

// go build && emlserver.exe -d \\ntsdtsg\c$\inetpub\mailroot\Drop

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"mime"
	"net/mail"
	"os"
	"path/filepath"
	"sort"
)

var (
	flagDropDir = flag.String("d", `\inetpub\mailroot\Drop`, "directory containing *.eml files")
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
	log.Println("serving *.eml files from", *flagDropDir)

	entries, err := ioutil.ReadDir(*flagDropDir)
	if err != nil {
		log.Fatalln(err)
	}
	// sort files desc
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].ModTime().Unix() > entries[j].ModTime().Unix()
	})
	for i, fi := range entries {
		log.Println(i, fi.ModTime(), fi.Name())
		emlfile := filepath.Join(*flagDropDir, fi.Name())
		err := parse(emlfile)
		if err != nil {
			log.Println(emlfile, err)
		}
	}
}

func parse(emlfile string) error {
	r, err := os.Open(emlfile)
	if err != nil {
		return err
	}
	m, err := mail.ReadMessage(r)
	if err != nil {
		return err
	}

	header := m.Header
	fmt.Println("Date:", header.Get("Date"))
	fmt.Println("From:", header.Get("From"))
	fmt.Println("To:", header.Get("To"))
	subj := header.Get("Subject")
	subj, err = worddecoder.DecodeHeader(subj)
	if err != nil {
		return fmt.Errorf("decoding subject: %v", err)
	}
	fmt.Println("Subject:", subj)

	cte := header.Get("Content-Transfer-Encoding")
	fmt.Println("Content-Transfer-Encoding:", cte)

	if cte == "base64" {
		dec := base64.NewDecoder(base64.StdEncoding, m.Body)
		body, err := ioutil.ReadAll(dec)
		if err != nil {
			return err
		}

		fmt.Printf("%s", body)
	} else {
		body, err := ioutil.ReadAll(m.Body)
		if err != nil {
			return err
		}

		fmt.Printf("%s", body)
	}

	return nil
}
