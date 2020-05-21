package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

type spin struct {
	mut    sync.Mutex
	locked int32
}

func (m *spin) lock(d time.Duration) bool {
	t0 := time.Now()
	var locked bool
	for !locked && time.Now().Sub(t0) < d {
		locked = m.tryLock()
	}
	return locked
}

func (m *spin) unlock() {
	atomic.StoreInt32(&m.locked, 0)
}

func (m *spin) tryLock() bool {
	return atomic.CompareAndSwapInt32(&m.locked, 0, 1)
}

func main() {
	s := spin{}
	fmt.Println("try to lock", s.tryLock())
	fmt.Println("try again", s.tryLock())
	fmt.Println("unlocked")
	s.unlock()
	fmt.Println("lock", s.lock(time.Second))
	fmt.Println("lock again", s.lock(time.Second))
}
