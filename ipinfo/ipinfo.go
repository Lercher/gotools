package main

import (
	"fmt"
	"log"
	"net"
)

// based on https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go

func main() {
	log.Println("This is ipinfo, (C) 2021 by Martin Lercher")

	ifaces, err := net.Interfaces()
	if err != nil {
		log.Fatal(err)
	}
	// handle err
	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			log.Panicln("if", i, err)
			continue
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				// fmt.Println("IPNet: ", "mask:", v.Mask.String(), "IP:", v.IP.String())
				if v.Mask.String() == "ffffffff" {
					fmt.Println("This is probably an IPv4 VPN connection:", v.IP.String())
				}
			case *net.IPAddr:
				// fmt.Println("IPAddr:", v.IP.String(), "zone:", v.Zone, "network:", v.Network())
			}
		}
	}
}
