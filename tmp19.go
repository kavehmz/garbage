package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		var i int64
		for 1 == 1 {
			i++
		}
	}()

	go func() {
		var i int64
		for 1 == 1 {
			i++
			time.Sleep(time.Millisecond * 100)
			fmt.Println(i)
		}
	}()

	time.Sleep(time.Second * 10)
}
