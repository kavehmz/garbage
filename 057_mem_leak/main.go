package main

import (
	"fmt"
	"runtime"
)

type t struct {
	a int
	b *[10000]int
}

var a []t

func main() {
	for i := 0; i < 10000; i++ {
		a = append(a, t{10, new([10000]int)})
	}
	noLeak(a)
	runtime.GC()
	mem()
	for i := 0; i < 10000; i++ {
		a = append(a, t{10, new([10000]int)})
	}
	leak(a)
	runtime.GC()
	mem()

}

func noLeak(a []t) {
	mem()
	for i := 1; i < 10000; i++ {
		a[i].b = nil
	}
	a = a[:1]

	runtime.GC()
	mem()
	fmt.Println("len", len(a))
}

func leak(a []t) {
	mem()
	for i := 1; i < 10000; i++ {
		a[i].b = nil
	}
	a = a[:1]

	runtime.GC()
	mem()
	fmt.Println("len", len(a))
}

func mem() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Alloc = %v MiB", mb(m.Alloc))
	fmt.Printf("\tTotalAlloc = %v MiB", mb(m.TotalAlloc))
	fmt.Printf("\tSys = %v MiB", mb(m.Sys))
	fmt.Printf("\tNumGC = %v\n", m.NumGC)
}

func mb(b uint64) uint64 {
	return b / 1024 / 1024
}
