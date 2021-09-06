package main

import (
	"fmt"
	"time"

	"github.com/pkg/profile"
)

const AN = 10000000
const BN = 1000

func byValue(a largeStruct) {
	b := [BN]int{}

	a.x = a.d[0]
	a.y = a.d[1]
	for i := 1; i < BN; i++ {
		b[i] = b[i] + a.x + a.y + i
	}
}

// checking if compile does any optimizatios
func byValueNotUseingTheLargeArray(a largeStruct) {
	b := [BN]int{}

	for i := 1; i < BN; i++ {
		b[i] = b[i] + a.x + a.y + i
	}
}

func byRef(a *largeStruct) {
	b := [BN]int{}

	for i := 1; i < BN; i++ {
		b[i] = b[i] + a.d[i-1] + i
	}
}

type largeStruct struct {
	d [AN]int
	x int
	y int
}

func load() {
	ls := largeStruct{}

	start := time.Now()
	byRef(&ls)
	fmt.Println("by ref first call:", time.Since(start))

	start = time.Now()
	byRef(&ls)
	fmt.Println("by ref second call:", time.Since(start))

	start = time.Now()
	byValue(ls)
	fmt.Println("by value first call:", time.Since(start))

	start = time.Now()
	byValue(ls)
	fmt.Println("by value second call:", time.Since(start))

	start = time.Now()
	byValueNotUseingTheLargeArray(ls)
	fmt.Println("by value no large use:", time.Since(start))

	start = time.Now()
	byValue(ls)
	fmt.Println("by value second call:", time.Since(start))

}

func main() {
	defer profile.Start(profile.MemProfile).Stop()

	start := time.Now()
	for time.Since(start) < 1*time.Second {
		load()
	}
}
