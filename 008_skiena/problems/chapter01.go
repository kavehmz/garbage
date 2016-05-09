package problems

// Divide with no / or *
func Divide(x, y int) int {
	if y > x {
		return 0
	}
	if y == x {
		return 1
	}
	return 1 + Divide(x-y, y)
}

// Bubble sort
func Bubble(a []int) {
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] > a[j] {
				a[i], a[j] = a[j], a[i]
			}
		}
	}
}

// Selection sort
func Selection(a []int) {
	for i := 0; i < len(a); i++ {
		k := i
		for j := i + 1; j < len(a); j++ {
			if a[k] > a[j] {
				k = j
			}
		}
		a[k], a[i] = a[i], a[k]
	}
}

// Insertion sort
func Insertion(a []int) {
	for i := 1; i < len(a); i++ {
		for k := i; k > 0 && a[k] < a[k-1]; k-- {
			a[k], a[k-1] = a[k-1], a[k]
		}
	}
}

// Merge sort
func Merge(a []int) {
	splitMerge(a, 0, len(a))
}

func splitMerge(a []int, begin, end int) {
	if end-begin < 2 {
		return
	}
	middle := (end + begin) / 2

	splitMerge(a, begin, middle)
	splitMerge(a, middle, end)
	b := make([]int, end-begin)

	i, j := begin, middle
	for k := begin; k < end; k++ {
		if i < middle && (j >= end || a[i] <= a[j]) {
			b[k-begin] = a[i]
			i++
		} else {
			b[k-begin] = a[j]
			j = j + 1
		}
	}
	for k := 0; k < end-begin; k++ {
		a[begin+k] = b[k]
	}
}

// Steps is
var Steps int

// Fib func to check the top down O(2^n) performance
func Fib(n int) int {
	Steps++
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 1
	}
	return Fib(n-1) + Fib(n-2)
}

// Quick func
func Quick(a []int) {
	quickSort(a, 0, len(a)-1)
}

func quickSort(a []int, lo, hi int) {
	if lo < hi {
		p := partition(a, lo, hi)
		quickSort(a, lo, p-1)
		quickSort(a, p+1, hi)
	}
}

func partition(a []int, lo, hi int) int {
	pivot := a[hi]
	i := lo

	for j := lo; j <= hi-1; j++ {
		if a[j] <= pivot {
			a[i], a[j] = a[j], a[i]
			i = i + 1
		}
	}
	a[i], a[hi] = a[hi], a[i]
	return i
}

// Heap sort
func Heap(a []int) {
	n := len(a)
	// Build heap (rearrange array)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(a, n, i)
	}
	// One by one extract an element from heap
	for i := n - 1; i >= 0; i-- {
		// Move current root to end
		a[0], a[i] = a[i], a[0]
		// call max heapify on the reduced heap
		heapify(a, i, 0)
	}
}

func heapify(a []int, n, i int) {
	largest := i
	l := 2*i + 1
	r := 2*i + 2

	if l < n && a[l] > a[largest] {
		largest = l
	}

	if r < n && a[r] > a[largest] {
		largest = r
	}

	if largest != i {
		a[i], a[largest] = a[largest], a[i]
		heapify(a, n, largest)
	}
}
