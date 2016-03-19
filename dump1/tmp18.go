package main

import "fmt"

func mapper(operator func(int) int, m []int) []int {
	for i, n := range m {
		m[i] = operator(n)
	}
	return m
}

func curriedMapper(operator func(int) int) func([]int) []int {
	return func(m []int) []int {
		for i, n := range m {
			m[i] = operator(n)
		}
		return m
	}
}

func main() {
	square := func(n int) int {
		return n * n
	}

	squaring := curriedMapper(square)
	fmt.Println(squaring([]int{1, 2, 3}))

	fmt.Println(mapper(square, []int{1, 2, 3}))
}
