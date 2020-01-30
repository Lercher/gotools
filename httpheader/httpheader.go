package main

import "log"
import "net/http"
import "os"

func main() {
	log.Println("This is httpheader, (C) 2020 by Martin Lercher")

	hdr := make(http.Header)

	k := ""
	for i, val := range os.Args {
		if i == 0 {
			continue
		}
		if i%2 == 1 {
			k = val
		} else {
			hdr.Add(k, val)
		}
	}
	// hdr.Add("umlauts-versal", "ÄÖÜ")
	// hdr.Add("umlauts-minuskel", "äöüß")
	// hdr.Add("colon-and-newline", "new-line:\nand-tab:\t\nand-backslash:\\//\\//\\//\\//\\ (it's cool man)")

	err := hdr.Write(os.Stdout)
	if err != nil {
		log.Fatal(err)

	}
}

/*
go run . multi a multi b multi cccc multi ddddddd
2020/01/30 16:24:26 This is httpheader, (C) 2020 by Martin Lercher
Colon-And-Newline: new-line: and-tab:    and-backslash:\//\//\//\//\ (it's cool man)
Multi: a
Multi: b
Multi: cccc
Multi: ddddddd
Umlauts-Minuskel: äöüß
Umlauts-Versal: ÄÖÜ

*/
