package main

import (
	"fmt"
	"math"
	"time"
)

type contract struct {
	price int
	ctype string
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
	for i := 0; i < 10000000; i++ {
		m = math.Exp(float64(i))
	}
	m++
	fmt.Println("loop time   :", float64(time.Now().UnixNano()-t)/1000000000)

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

}
