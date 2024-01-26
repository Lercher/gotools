package main

// go run main.go environment.go sample.txt sample.json

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	log.Println("This is", os.Args[0], "(C) 2018 by Martin Lercher")

	jsonname := os.Args[2]
	text, err := ioutil.ReadFile(jsonname)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print("Loaded datafile", jsonname)

	m := make(map[string]interface{})
	err = json.Unmarshal(text, &m)
	if err != nil {
		log.Fatalln(err)
	}
	log.Print("Parsed data", m)

	templatename := os.Args[1]
	var e *Environment
	rawtpl := template.Must(template.New("").Funcs(e.FuncMap()).ParseFiles(templatename))
	log.Println("Loaded template", templatename, rawtpl.DefinedTemplates())

	e = NewEnvironment()
	log.Println("Created new render environment:", *e)
	tpl := template.Must(rawtpl.Clone()).Funcs(e.FuncMap())
	err = tpl.ExecuteTemplate(os.Stdout, templatename, m)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Bound values:")
	for _, v := range e.BoundValue {
		log.Println(" ", v)
	}

	for k, v := range e.BoundValue {
		log.Println("Modifying first bound value", k)
		nv := "New value"
		log.Print(v.Property, " from ", v.Value(), " to ", nv)
		err = v.Put(nv)
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("->", v.Value())
		break
	}
	log.Println(m)
}
