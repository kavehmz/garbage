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
