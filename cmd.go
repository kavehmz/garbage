package main

import (
	"bufio"
	"fmt"
	"log"
	"os/exec"
	"strconv"
	"time"
)

func price(n int) {
	var (
		err error
	)

	cmdName := "/tmp/price.pl"
	cmdArgs := []string{strconv.Itoa(n)}
	cmd := exec.Command(cmdName, cmdArgs...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Println("error: ", err)
	}

	if err := cmd.Start(); err != nil {
		fmt.Println("error: ", err)
	}

	in := bufio.NewScanner(stdout)

	for in.Scan() {
		log.Printf(in.Text())
	}
	if err := in.Err(); err != nil {
		fmt.Println("error: ", err)
	}
}

func main() {
	for i := 0; i < 150; i++ {
		fmt.Println("start", i)
		go price(i)
	}
	time.Sleep(100 * time.Second)
}
