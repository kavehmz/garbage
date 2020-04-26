// Lomuto partition scheme
package main

import (
	"fmt"
)

/*

algorithm quicksort(A, lo, hi) is
    if lo < hi then
        p := partition(A, lo, hi)
        quicksort(A, lo, p - 1)
        quicksort(A, p + 1, hi)

algorithm partition(A, lo, hi) is
    pivot := A[hi]
    i := lo
    for j := lo to hi do
        if A[j] < pivot then
            swap A[i] with A[j]
            i := i + 1
    swap A[i] with A[hi]
	return i
*/

func main() {
	a := []int{2, 3, 4, 5, 6, 31, 2, 1, 23}
	fmt.Println(a)
	quicksort(a, 0, 8)
	fmt.Println(a)
}

func quicksort(A []int, lo, hi int) {
	if lo <= hi {
		p := partition(A, lo, hi)
		quicksort(A, lo, p-1)
		quicksort(A, p+1, hi)
	}

}

func partition(A []int, lo, hi int) int {
	pivot := A[hi]
	i := lo
	for j := lo; j <= hi; j++ {
		if A[j] < pivot {
			A[i], A[j] = A[j], A[i]
			i = i + 1
		}
	}
	A[i], A[hi] = A[hi], A[i]
	return i
}
