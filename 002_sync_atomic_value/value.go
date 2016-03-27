package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type timeConfig struct {
	update time.Time
}

var config atomic.Value

func showConfig() {
	c := config.Load().(timeConfig)
	fmt.Println(c.update)
}

func main() {
	c := timeConfig{update: time.Now()}
	config.Store(c)
	fmt.Println(c.update)
	time.Sleep(100 * time.Millisecond)
	go showConfig()
	c.update = time.Now()
	config.Store(c)
	time.Sleep(100 * time.Millisecond)
}
