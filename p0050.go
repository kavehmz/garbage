package primes

import "fmt"
import "math"

const (
	Max = 1000000
)

var primes []int

func is_prime(n int) bool {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func find_primes() []int {
	var p []int
	for i := 2; i < Max; i++ {
		if is_prime(i) {
			p = append(p, i)
		}
	}
	return p
}

func main() {
	fmt.Println(len(find_primes()))
}
