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
