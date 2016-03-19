package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"

	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
)

func main() {

	sock, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK"))
	if err != nil {
		log.Fatal(err)
	}

	agent := agent.NewClient(sock)

	signers, err := agent.Signers()
	if err != nil {
		log.Fatal(err)
	}

	// An SSH client is represented with a ClientConn. Currently only
	// the "password" authentication method is supported.
	//
	// To authenticate with the remote server you must pass at least one
	// implementation of AuthMethod via the Auth field in ClientConfig.
	auths := []ssh.AuthMethod{ssh.PublicKeys(signers...)}
	config := &ssh.ClientConfig{
		User: "root",
		Auth: auths,
	}

	//	config := &ssh.ClientConfig{
	//		User: "username",
	//		Auth: []ssh.AuthMethod{
	//			ssh.Password("yourpassword"),
	//		},
	//	}

	client, err := ssh.Dial("tcp", "128.199.236.127:22", config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}

	// Each ClientConn can support multiple interactive sessions,
	// represented by a Session.
	session, err := client.NewSession()
	if err != nil {
		panic("Failed to create session: " + err.Error())
	}
	defer session.Close()

	out, err := session.StdoutPipe()
	if err != nil {
		panic("Failed to create session pipe: " + err.Error())
	}

	scanner := bufio.NewScanner(out)

	go func() {
		for scanner.Scan() {
			fmt.Printf("new line received: %s\n", scanner.Text())
		}
	}()

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	//if err := session.Run("while [ 1 ];do date;sleep 1;done"); err != nil {
	if err := session.Run("`"); err != nil {
		panic("Failed to run: " + err.Error())
	}

	session.Wait()
}
