package main

import (
	"fmt"
	"math"
	"strconv"
	"time"
)

type contract struct {
	price int
	ctype string
}

const sqrt2PI = 2.506628274631

var t = strconv.Itoa(5)

func pdf(x, m, s float64) float64 {
	if s < 0 {
		panic("die")
	}
	z := (x - m) / s
	return math.Exp(-0.5*z*z) / (sqrt2PI * s)
}

func main() {
	t := time.Now().UnixNano()
	n := 0
	for i := 0; i < 10000000; i++ {
		n++
	}
	fmt.Println("loop time   :", float64(time.Now().UnixNano()-t)/1000000000)

	t = time.Now().UnixNano()
	m := 0.0
	for f := 0.0; f < 10000000; f++ {
		m = math.Exp(f)
	}
	m++
	fmt.Println("exp time   :", float64(time.Now().UnixNano()-t)/1000000000)

	t = time.Now().UnixNano()
	for i := 0; i < 10000; i++ {
		a := make([]int, 1000)
		// fmt.Println(a[10])
		for j := 0; j < 10000; j++ {
			a = append(a, j)
		}
	}
	fmt.Println("append time :", float64(time.Now().UnixNano()-t)/1000000000)

	t = time.Now().UnixNano()
	var p contract
	for i := 0; i < 10000000; i++ {
		p = contract{1, "some type"}
	}
	p.ctype = "r"
	fmt.Println("struct time   :", float64(time.Now().UnixNano()-t)/1000000000)

	t = time.Now().UnixNano()
	for i := 0; i < 10000000; i++ {
		pdf(1.2, 1.3, 1.4)
	}
	fmt.Println("pdf time   :", float64(time.Now().UnixNano()-t)/1000000000)

}
