package main

import (
	"fmt"
	"time"
)

const (
	windowSize = 200000
	msgCount   = 1000000
)

type message []byte

type channel [windowSize]message

var worst time.Duration
var average int64

func mkMessage(n int) message {
	m := make(message, 1024)
	for i := range m {
		m[i] = byte(n)
	}
	return m
}

func pushMsg(c *channel, highID int) {
	start := time.Now()
	m := mkMessage(highID)
	(*c)[highID%windowSize] = m
	elapsed := time.Since(start)
	average += elapsed.Nanoseconds()
	if elapsed > worst {
		worst = elapsed
	}
}

func main() {
	var c channel
	for i := 0; i < msgCount; i++ {
		pushMsg(&c, i)
	}
	fmt.Println("Worst push time: ", worst)
	fmt.Println("averg push time: ", float64(average)/(msgCount))
}
