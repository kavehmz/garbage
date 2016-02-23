package main

import "fmt"

type typeFunc func(n int) int

func (f typeFunc) centimeter(n int) int {
	n = n * 100
	return f(n)
}

func (f typeFunc) meter(n int) int {
	n = n * 1
	return f(n)
}

func main() {
	var cubicvolume typeFunc = func(n int) int {
		return n * n * n
	}

	fmt.Println(cubicvolume(5))
	fmt.Println(cubicvolume.meter(5))
	fmt.Println(cubicvolume.centimeter(5))
}
