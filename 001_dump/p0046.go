package main

import "fmt"

func is_composite(n int) bool {
	for i := 3; i < n; i = i + 2 {
		if n%i == 0 {
			return true
		}
	}
	return false
}

func is_twice_a_square(n int, p []int) bool {
	pc := 0
	for p[pc] != 0 && p[pc] < n {
		for i := 1; i < n; i++ {
			if n == p[pc]+2*i*i {
				return true
			}
		}
		pc++
	}
	return false
}

func main() {
	max := 10000
	primes := make([]int, max)
	pc := 0
	for n := 3; n < max; n = n + 2 {
		if is_composite(n) {
			if !is_twice_a_square(n, primes) {
				fmt.Println(n)
			}
		} else {
			primes[pc] = n
			pc++
		}
	}
}
