package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

const (
	windowSize = 200000
	msgCount   = 10000000
)

type message []byte

type channel [windowSize]message

var worst time.Duration

func mkMessage(n int) message {
	m := make(message, 1024)
	for i := range m {
		m[i] = byte(n)
	}
	return m
}

var pool chan bool
var l sync.Mutex

func pushMsg(c *channel, highID int) {
	start := time.Now()
	m := mkMessage(highID)
	(*c)[highID%windowSize] = m
	elapsed := time.Since(start)
	l.Lock()
	if elapsed > worst {
		worst = elapsed
	}
	l.Unlock()
	<-pool
}

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	pool = make(chan bool, 4)
	var c channel
	for i := 0; i < msgCount; i++ {
		pool <- true
		go pushMsg(&c, i)
	}
	for i := 0; i < 4; i++ {
		pool <- true
	}
	fmt.Println("Worst push time: ", worst)
}
