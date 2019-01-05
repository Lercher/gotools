package main

// go get && go run main.go
// see https://godoc.org/github.com/go-stomp/stomp

import (
	"log"

	"github.com/jjeffery/stomp"
)

func main() {
	con, err := stomp.Dial("tcp", "item-s69570:61613")
	if err != nil {
		log.Fatalln("dial:", err)
	}
	defer con.Disconnect()

	log.Println("STOMP version", con.Version())
	log.Println("connected to", con.Server())
	log.Println("session", con.Session())

	err = con.Send("SampleQueue", "string", []byte("lorem ipsum"), stomp.SendOpt.Receipt)
	if err != nil {
		log.Fatalln("send:", err)
	}

	err = con.Disconnect()
	if err != nil {
		log.Fatalln("disconnect:", err)
	}
}
