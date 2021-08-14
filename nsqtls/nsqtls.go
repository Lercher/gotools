package main

// NSQ - https://nsq.io/

import (
	"crypto/tls"
	"crypto/x509"
	_ "embed"
	"encoding/hex"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	nsq "github.com/nsqio/go-nsq"
)

//embed:ca.cert
var cacertPEM []byte

var (
	flagTopic   = flag.String("topic", "testtopic", "topic to consume")
	flagChannel = flag.String("channel", "consolePeeker#ephemeral", "channel name (#ephemeral marks a temporary channel)")
	flagNSQD    = flag.String("nsqd", "linux-pm81:4150", "host:port of the nsqd service")
	flagKey     = flag.String("key", "client.key", "TLS private key file in PEM format")
	flagCert    = flag.String("cert", "client.cert", "TLS public key file in PEM format")
)

type handler string

func (h handler) HandleMessage(m *nsq.Message) error {
	log.Println(h)
	fmt.Println(string(m.Body))
	return nil
}

func verifyCert(rawCerts [][]byte, verifiedChains [][]*x509.Certificate) error {
	for i := range rawCerts {
		fmt.Println("----------------- raw cert", i, "------------------------")
		fmt.Println(hex.Dump(rawCerts[i]))
		// TODO: verify this cert!
	}
	return nil
}

func main() {
	log.Println("This is", os.Args[0], "(C) 2021 by Martin Lercher")
	flag.Parse()

	cfg := nsq.NewConfig()
	cfg.UserAgent = os.Args[0]

	// using TLS, see also https://groups.google.com/g/nsq-users/c/PkkPN9z7gPc
	cfg.TlsV1 = true
	cert, err := tls.LoadX509KeyPair(*flagCert, *flagKey)
	if err != nil {
		log.Fatalln("loading TLS key/cert files:", err)
	}
	cfg.TlsConfig = &tls.Config{
		Certificates:          []tls.Certificate{cert},
		VerifyPeerCertificate: verifyCert,
		InsecureSkipVerify:    true,
	}

	consumer, err := nsq.NewConsumer(*flagTopic, *flagChannel, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	consumer.AddHandler(handler(*flagTopic))

	log.Println("connecting to NSQD", *flagNSQD, "and subscribing to topic", *flagTopic, "on channel", *flagChannel, "...")
	err = consumer.ConnectToNSQD(*flagNSQD)
	if err != nil {
		log.Fatalln(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Println("connecting to NSQD", *flagNSQD, "and producing to to topic", *flagTopic, "...")
		prod, err := nsq.NewProducer(*flagNSQD, cfg)
		if err != nil {
			log.Println("Sorry, can't send:", err)
			return
		}
		defer prod.Stop()

		for i := 0; i < 4; i++ {
			msg := fmt.Sprint("message current time is: ", time.Now().Local().Format(time.Stamp))
			prod.Publish(*flagTopic, []byte(msg))
			time.Sleep(time.Second)
		}
		log.Println("all messages produced, hit ^C to stop")
	}()
	<-sigChan
	log.Println("stoping")

	consumer.Stop()
	<-consumer.StopChan
	log.Println("shutdown complete")
}
