package main

import (
	"fmt"
	"math"
)

type benchTest struct {
	value float64
}

func bench(f float64) float64 {
	var s benchTest
	s.value = math.Sin(f)
	return s.value
}

func main() {
	tt := math.Sin(1.0)
	for i := 0; i < 10000000; i++ {
		tt = bench(tt)
	}
	fmt.Println(tt)
}
