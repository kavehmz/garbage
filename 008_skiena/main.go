package main

import (
	"fmt"

	"github.com/kavehmz/garbage/008_skiena/problems"
)

func main() {
	fmt.Println(problems.Divide(49, 7))

	a := []int{1, 4, 3, 1, 7, 9, 2, 8}
	problems.Merge(a)
	fmt.Println(a)
}
