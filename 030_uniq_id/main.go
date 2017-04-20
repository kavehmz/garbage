package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"time"
)

type UID struct {
	time  uint32
	count uint32
	gid   uint64
}

type ID struct {
	gid   uint64
	count uint32
	last  uint32
}

func (i *ID) next() UID {
	now := uint32(time.Now().Unix())
	if now > i.last {
		i.count = 0
	}
	i.count++
	return UID{now, i.count, i.gid}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	id := ID{gid: rand.Uint64()}
	t := time.Now().UnixNano()
	for i := 0; i < 50*1000000; i++ {
		id.next()
	}
	fmt.Println("seconds it took:", float64(time.Now().UnixNano()-t)/1000000000)
	runtime.CPUProfile()
