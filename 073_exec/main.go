package main

import (
	"fmt"
	"syscall"
)

func main() {
	// Checking how can I exce and control the env vars
	err := syscall.Exec("./export.sh", []string{"", "ARG1"}, []string{"FOO=BAR", "FOO=BAR & MAR"})

	if err != nil {
		fmt.Println(err)
	}
}
