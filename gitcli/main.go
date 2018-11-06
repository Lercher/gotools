package main

import (
	_ "fmt"
	"log"

	"gopkg.in/src-d/go-git.v4"
)

func main() {
	log.Println("This is gitcli, to test a go based git lib")

	r, err := git.PlainOpen("..")
	if err != nil {
		log.Fatal(err)
	}

	origin, err := r.Remote("origin")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("String\t", origin.String())
	log.Println("URLs[0]\t", origin.Config().URLs[0])
}
