package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/mail"
	"os"
	"path/filepath"
	"sort"
)

type server struct {
	rootpath  string
	directory string
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	mb, err := s.list(id)
	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Fprintln(w, "OK", mb)
}

func (s *server) prefix() string {
	return fmt.Sprintf("/%s", s.rootpath)
}

func (s *server) list(id string) (*mailbox, error) {
	entries, err := ioutil.ReadDir(*flagDropDir)
	if err != nil {
		log.Fatalln(err)
	}
	// sort files desc
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].ModTime().Unix() > entries[j].ModTime().Unix()
	})
	mb := &mailbox{}
	for _, fi := range entries {
		emlfile := filepath.Join(*flagDropDir, fi.Name())
		item, err := parse(emlfile, id)
		if err != nil {
			return nil, err
		}
		if item.ID == id {
			mb.Selected = item
		}
		mb.All = append(mb.All, item)
	}
	return mb, nil
}

func parse(emlfile, id string) (*item, error) {
	r, err := os.Open(emlfile)
	if err != nil {
		return nil, err
	}
	defer r.Close()

	m, err := mail.ReadMessage(r)
	if err != nil {
		return nil, fmt.Errorf("%s: %v", emlfile, err)
	}

	item := &item{}
	item.ID = filepath.Base(emlfile)
	header := m.Header
	item.From = header.Get("From")
	item.Date = header.Get("Date")
	item.To = header.Get("To")
	subj := header.Get("Subject")
	subj, err = worddecoder.DecodeHeader(subj)
	if err != nil {
		return nil, fmt.Errorf("decoding subject: %v", err)
	}
	item.Subject = subj

	if item.ID == id {
		cte := header.Get("Content-Transfer-Encoding")
		if cte == "base64" {
			dec := base64.NewDecoder(base64.StdEncoding, m.Body)
			body, err := ioutil.ReadAll(dec)
			if err != nil {
				return nil, fmt.Errorf("decoding body: %v", err)
			}
			item.Body = string(body)
		} else {
			body, err := ioutil.ReadAll(m.Body)
			if err != nil {
				return nil, fmt.Errorf("reading body: %v", err)
			}
			item.Body = string(body)
		}
	}

	return item, nil
}
