package main

import (
	"io"
	"log"
	"os"
	"strings"
	"text/template"
)

type data struct {
	W io.Writer
	Text string
	T *template.Template
}


func main() {
	log.Println("This os template in template, (C) 2020 by Martin Lercher")

	t := template.Must(template.ParseGlob("*.txt"))
	
	data :=data{W: os.Stdout}
	for _, tt := range t.Templates() {
		log.Println("template", tt.Name())
		if strings.HasPrefix(tt.Name(), "inner") {
			data.Text = tt.Name()
			data.T = tt
		}
	}

	err := t.ExecuteTemplate(data.W, "master", data)
	if err != nil {
		log.Fatal(err)
	}
}