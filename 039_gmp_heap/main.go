package main

import (
	"fmt"
	"log"
	"runtime"
	"time"
)

func runAndWait(next, wait chan bool) {
	next <- true
	<-wait
}

func startAll() {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Println(mem.Alloc)
	log.Println(mem.TotalAlloc)
	log.Println(mem.HeapAlloc)
	log.Println(mem.HeapSys)

	next := make(chan bool)
	wait := make(chan bool)
	for i := 0; i <= 1000000; i++ {
		go runAndWait(next, wait)
		<-next
	}
	fmt.Println(runtime.NumGoroutine())

	runtime.GC()

	runtime.ReadMemStats(&mem)
	log.Println(mem.Alloc)
	log.Println(mem.TotalAlloc)
	log.Println(mem.HeapAlloc)
	log.Println(mem.HeapSys)

	close(wait)
	time.Sleep(time.Second * 5)
}

func main() {
	startAll()
	fmt.Println(runtime.NumGoroutine())
	runtime.GC()

	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	log.Printf("%+v\n", mem)

}
