package main

// go get && go run main.go

// see https://godoc.org/github.com/go-stomp/stomp
// and http://activemq.apache.org/stomp.html

import (
	"log"
	"time"

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

const q = "Sample.Queue"
const host = "item-s69570:61613"

func main() {
	con, err := stomp.Dial(
		"tcp", host,
		stomp.ConnOpt.Header("client-id", "hostname.sample"),
	)
	if err != nil {
		log.Fatalln("dial:", err)
	}
	defer con.Disconnect()

	log.Println("STOMP version", con.Version())
	log.Println("connected to", con.Server())
	log.Println("session", con.Session())
	log.Println()

	err = con.Send(
		q,
		"string",
		[]byte("lorem ipsum"),
		stomp.SendOpt.Receipt,
		stomp.SendOpt.Header("persistent", "true"),
		stomp.SendOpt.Header("correlation-id", "+49897482400"),
	)
	if err != nil {
		log.Fatalln("send:", err)
	}

	// now receive something:
	sub, err := con.Subscribe(
		q,
		stomp.AckClient,
		stomp.SubscribeOpt.Header("activemq.prefetchSize", "1"),
	)
	if err != nil {
		log.Fatalln("subscribe:", err)
	}

loop:
	for {
		select {
		case msg := <-sub.C:
			if msg.Err != nil {
				log.Println("msg:", msg.Err)
				break
			}

			l := msg.Header.Len()
			for i := 0; i < l; i++ {
				k, v := msg.Header.GetAt(i)
				log.Println(k, v)
			}
			log.Println("content-type:", msg.ContentType)
			log.Println("destination: ", msg.Destination)
			log.Println("content:     ", string(msg.Body))
			log.Println()

			err = con.Ack(msg)
			if msg.Err != nil {
				log.Println("ack:", msg.Err)
				break
			}
		case <-time.After(3 * time.Second):
			break loop
		}
	}
	err = sub.Unsubscribe()
	if err != nil {
		log.Fatalln("unsubscribe:", err)
	}

	err = con.Disconnect()
	if err != nil {
		log.Fatalln("disconnect:", err)
	}
	log.Println("disconnected")
}
