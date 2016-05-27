package main

import "fmt"

func main() {

	a := []int{1, 2, 3, 4, 5, 6, 7, 8}
	// We cut a but values are still there! Imagine if these values were pointer to larger strcutures and we have a possiblity of memory leakage
	a = append(a[:1], a[3:]...)

	c := a[:8]

	// To see that values are still there see how c after assignment operation still have last a values
	// a lenght has changed but its cap did had not changed.
	fmt.Println(a)
	fmt.Println(c)
}
