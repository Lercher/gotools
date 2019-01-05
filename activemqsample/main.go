package main

// go get && go run main.go

// see https://godoc.org/github.com/go-stomp/stomp
// and http://activemq.apache.org/stomp.html

import (
	"log"

	"github.com/go-stomp/stomp"
)

/* SEND, STOMP Header, Description

correlation-id
Good consumers will add this header to any responses they send.

expires
Expiration time of the message.

JMSXGroupID
Specifies the Message Groups.

JMSXGroupSeq
Optional header that specifies the sequence number in the Message Groups.

persistent
Whether or not the message is persistent.

priority
Priority on the message.

reply-to
Destination you should send replies to.

type
Type of the message. [need not be a header, it's the 2nd send parameter]
*/

func main() {
	con, err := stomp.Dial(
		"tcp", "item-s69570:61613", 
		stomp.ConnOpt.Header("client-id", "hostname.sample"),
	)
	if err != nil {
		log.Fatalln("dial:", err)
	}
	defer con.Disconnect()

	log.Println("STOMP version", con.Version())
	log.Println("connected to", con.Server())
	log.Println("session", con.Session())

	err = con.Send(
		"SampleQueue",
		"string",
		[]byte("lorem ipsum"),
		stomp.SendOpt.Receipt,
		stomp.SendOpt.Header("persistent", "true"),
		stomp.SendOpt.Header("correlation-id", "+49897482400"),
	)
	if err != nil {
		log.Fatalln("send:", err)
	}

	err = con.Disconnect()
	if err != nil {
		log.Fatalln("disconnect:", err)
	}
}
