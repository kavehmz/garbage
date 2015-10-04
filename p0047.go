package main

import "fmt"

func approx_n_power(n uint64) uint64 {
	var s uint64 = 1
	for i := uint64(1); i <= n; i++ {
		s = (s * n) % 10000000000
	}
	return s
}

func main() {
	var s uint64
	for i := 1; i <= 1000; i++ {
		s += approx_n_power(uint64(i))
	}
	fmt.Println(s)
}
