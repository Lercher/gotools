package main

// see also: https://github.com/nsqio/nsq/blob/master/apps/nsq_tail/nsq_tail.go and https://nsq.io/overview/quick_start.html

// go get && go run main.go -topic foo.bar
// curl -d 'hello world 3' 'http://127.0.0.1:4151/pub?topic=foo.bar'
// http://127.0.0.1:4171/

//// NSQ - https://nsq.io/
// cd ~/nsq-1.1.0.linux-amd64.go1.10.3/bin/
// ./nsqlookupd 
// ./nsqd --lookupd-tcp-address=127.0.0.1:4160
// ./nsqadmin --lookupd-http-address=127.0.0.1:4161

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	nsq "github.com/nsqio/go-nsq"
)

var (
	flagTopic     = flag.String("topic", "", "topic to consume")
	flagChannel   = flag.String("channel", "consolePeeker#ephemeral", "channel name (#ephemeral marks a temporary channel)")
	flagDirectory = flag.String("dir", "localhost:4161", "host:port of the nsqlookupd directory services http port")
)

type handler string

func (h handler) HandleMessage(m *nsq.Message) error {
	log.Println(h)
	fmt.Println(string(m.Body))
	return nil
}

func main() {
	log.Println("This is", os.Args[0], "(C) 2019 by Martin Lercher")
	flag.Parse()

	cfg := nsq.NewConfig()
	cfg.UserAgent = os.Args[0]

	consumer, err := nsq.NewConsumer(*flagTopic, *flagChannel, cfg)
	if err != nil {
		log.Fatalln(err)
	}

	consumer.AddHandler(handler(*flagTopic))

	log.Println("connecting to NSQLookupd", *flagDirectory, "and subscribing to topic", *flagTopic, "on channel", *flagChannel, "...")
	err = consumer.ConnectToNSQLookupd(*flagDirectory)
	if err != nil {
		log.Fatalln(err)
	}

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	<-sigChan
	log.Println("stoping")

	consumer.Stop()
	<-consumer.StopChan
	log.Println("shutdown complete")
}
