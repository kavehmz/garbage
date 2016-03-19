package main

import "fmt"

func myFor(i, n int, f func(int)) {
	if i >= n {
		return
	}
	i++
	f(i)
	myFor(i, n, f)
}

func main() {

	i := 0
	myFor(i, 1000000, func(i int) {
		fmt.Println(i)
	})
}
