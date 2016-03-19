//usr/bin/env go run $0 $@;exit

package main

import (
	"os"
	"time"
)

func init() { print("A\n") }
func f() {
	defer func() {
		print("on exit\n")
	}()
	print("bye\n")
}

func init() { print("B\n") }
func main() {
	go f()

	time.Sleep(time.Second)
	os.Exit(0)
}
