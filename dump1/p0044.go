package main

import "math"
import "fmt"

func is_p(p int) bool {
	n := (1 + math.Sqrt(1+24*float64(p))) / 6
	return math.Trunc(n) == n
}

func p(n int) int {
	return n * (3*n - 1) / 2
}

func main() {
	for i := 2; i < 10000; i++ {
		for j := i - 1; j >= 1; j-- {
			if is_p(p(i)+p(j)) && is_p(p(i)-p(j)) {
				fmt.Println(i, p(i)-p(j))
				return
			}
		}

	}
}
