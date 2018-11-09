package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
)

var (
	flagPass = flag.String("pass", "", "password for ssh")
)

func main() {
	log.Println("This is sshtest: testing ssh and sftp")
	flag.Parse()

	if *flagPass == "" {
		flag.Usage()
		log.Fatal("no password given")
	}
	config := &ssh.ClientConfig{
		User: "lercher",
		Auth: []ssh.AuthMethod{
			ssh.Password(*flagPass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		BannerCallback:  ssh.BannerDisplayStderr(),
		Timeout:         time.Second * 2,
	}

	// Dial the ssh server:
	con, err := ssh.Dial("tcp", "localhost:22", config)
	if err != nil {
		log.Fatal("unable to connect: ", err)
	}
	defer con.Close()

	log.Println("ssh client version", string(con.ClientVersion()))

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := con.NewSession()
	if err != nil {
		log.Fatal("Failed to create ssh session: ", err)
	}
	defer session.Close()

	// Once a Session is created, you can execute a single command on
	// the remote side using the Run method.
	session.Stdout = os.Stdout
	cmd := `pwd`
	log.Println(cmd)
	if err := session.Run(cmd); err != nil {
		log.Fatal("Failed to run", cmd, ":", err)
	}
	log.Println("completed", cmd)

	// now upload a simple file
	sh, err := sftp.NewClient(con)
	if err != nil {
		log.Fatal("opening sftp:", err)
	}
	fn := "hello_world_sftp.txt"
	f, err := sh.Create(fn)
	if err != nil {
		log.Fatal(err)
	}
	n, err := fmt.Fprintf(f, "Hello sftp world, it's %v\nAnd here are some umlauts: äöüÄÖÜß\n", time.Now())
	if err != nil {
		log.Fatal("writing", fn, ":", err)
	}
	log.Println("written", n, "bytes")
}
