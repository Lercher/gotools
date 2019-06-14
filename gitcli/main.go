package main

import (
	_ "fmt"
	"log"
	"path/filepath"
	"io/ioutil"
	"os"

	"gopkg.in/src-d/go-git.v4"
	"gopkg.in/src-d/go-git.v4/plumbing/transport/ssh"

	cryptossh  "golang.org/x/crypto/ssh"
)

func main() {
	log.Println("This is gitcli, to test a go based git lib")

	r, err := git.PlainOpen("..")
	if err != nil {
		log.Fatal(err)
	}

	origin, err := r.Remote("origin")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("String\t", origin.String())
	log.Println("URLs[0]\t", origin.Config().URLs[0])
	log.Println("Config().Name\t", origin.Config().Name)

	// login to git server using ssh private key
	// avoids this error message, if no Auth is given at git.FetchOptions:
	//
	//   error creating SSH agent: "SSH agent requested but Pageant not running"
	//
	// credits to:
	// https://github.com/wstrange at https://github.com/src-d/go-git/issues/550#issuecomment-323182885
	//
	s := filepath.Join(os.Getenv("USERPROFILE"), ".ssh", "id_rsa")
	sshKey, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatalf("readfile: %v", err)
	}
	signer, err := cryptossh.ParsePrivateKey(sshKey)
	if err != nil {
		log.Fatalf("parseprivatekey: %v", err)
	}
	gitAuth := &ssh.PublicKeys{User: "git", Signer: signer}
	log.Println(gitAuth)

	log.Println("fetch", origin.Config().Name)
	err = r.Fetch(&git.FetchOptions{
		Auth:       gitAuth,
		RemoteName: origin.Config().Name,
	})
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DONE")
}
