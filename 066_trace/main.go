package main

import (
	"log"
	"os"
	"runtime/trace"
	"sync"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		log.Panic(err)
	}
	defer trace.Stop()

	var wg sync.WaitGroup
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			t := 0
			for i := 0; i < 3*1e8; i++ {
				t += 2
			}
			log.Println("Done", t)
		}()

	}
	wg.Wait()
}
