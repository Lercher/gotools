package main

// go get && go run main.go

import (
	"log"

	"github.com/jjeffery/stomp"
)

func main() {
	conn, err := stomp.Dial("tcp", "item-s69570:61613")
	if err != nil {
		log.Fatalln("dial:", err)
	}
	log.Println("STOMP version", conn.Version())
	log.Println("connected to", conn.Server())
	log.Println("session", conn.Session())

	err = conn.Send("SampleQueue", "string", []byte("lorem ipsum"), stomp.SendOpt.Receipt)
	if err != nil {
		log.Fatalln("send:", err)
	}

	err = conn.Disconnect()
	if err != nil {
		log.Fatalln("disconnect:", err)
	}
}
