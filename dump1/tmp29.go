package main

import (
	"fmt"
	"sync"
	"time"
)

type person struct {
	name      string
	timestamp time.Time
}

func main() {
	var pool sync.Pool
	pool.New = func() interface{} {
		return person{name: "name", timestamp: time.Now()}
	}
	var p, q, r person

	p = pool.Get().(person)
	fmt.Println(p.timestamp.String())
	pool.Put(p)
	r = pool.Get().(person)
	q = pool.Get().(person)
	fmt.Println(r.timestamp.String())
	fmt.Println(q.timestamp.String())

}
