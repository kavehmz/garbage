package main

import "fmt"
import "strconv"
import "sort"

var sum int64

func is_prime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

var pc map[string][]int = make(map[string][]int)

func permutaion_count(n int) {
	var chars []rune = []rune(strconv.Itoa(n))
	var digits []int = make([]int, len(chars))
	for key, value := range chars {
		digits[key] = int(value)
	}
	sort.Sort(sort.IntSlice(digits))
	for key, value := range digits {
		chars[key] = rune(value)
	}
	pc[string(chars)] = append(pc[string(chars)], n)
}

func main() {
	for i := 1000; i < 9999; i++ {
		if is_prime(i) {
			permutaion_count(i)
		}
	}

	for _, v := range pc {
		for a := 0; a < len(v); a++ {
			for b := a + 1; b < len(v); b++ {
				for c := b + 1; c < len(v); c++ {
					if v[b]-v[a] == v[c]-v[b] {
						fmt.Println(v[a], v[b], v[c])
					}
				}
			}
		}
	}

}
