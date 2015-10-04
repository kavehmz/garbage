package main

import "math"
import "fmt"

func is_hexagonal(p int) bool {
	n := (1 + math.Sqrt(1+8*float64(p))) / 4
	return math.Trunc(n) == n
}

func is_pentagonal(p int) bool {
	n := (1 + math.Sqrt(1+24*float64(p))) / 6
	return math.Trunc(n) == n
}

func is_triangle(x int) bool {
	n := (-1 + math.Sqrt(1+8*float64(x))) / 2
	return math.Trunc(n) == n
}

func hexagonal(n int) int {
	return n * (2*n - 1)
}

func main() {
	i := 144
	h := hexagonal(i)
	for ; !is_hexagonal(h) || !is_pentagonal(h) || !is_triangle(h); i++ {
		h = hexagonal(i)
	}

	fmt.Println(h)
}
