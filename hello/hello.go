package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
	"time"
)

func server() {
	// http.Handle("/foo", fooHandler)

	mux := http.NewServeMux()
	mux.HandleFunc("/foo/", func(w http.ResponseWriter, r *http.Request) {
		// http://localhost:8080/foo/something?q=1&s=2
		fmt.Fprintf(w, 
			"/foo/* handler here, %q, query %q, q=%q s=%q", 
			html.EscapeString(r.URL.Path), 
			html.EscapeString(r.URL.RawQuery), 
			html.EscapeString(r.URL.Query()["q"][0]),
			html.EscapeString(r.URL.Query()["s"][0]),
		)
	})

	mux.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "/bar only handler here, %q", html.EscapeString(r.URL.Path))
	})

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "I am the catch all, %q", html.EscapeString(r.URL.Path))
	})

	fmt.Println(time.Now(), "ListenAndServe 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
	fmt.Println(time.Now(), "after ListenAndServe")
}

func main() {
	fmt.Println("hello, world. Press Enter to exit.")
	go server()
	var input string
	fmt.Scanln(&input)
	fmt.Println("exit")
}
