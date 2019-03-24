package main

import (
	"log"
	"math"
)

func main() {
	log.Println(maxSqrt(2, 100000000000))
}

func maxSqrt(a, b int64) int64 {
	max := int64(0)
	number := int64(0)
	for i := int64(2); i <= int64(math.Sqrt(float64(b))); i++ {
		count := int64(0)
		for n := i * i; n < b && n > 1; n = n * n {
			count++
			if n > a && count > max {
				number = i
				max = count
			}
		}
	}
	log.Println("the base number which can be multiplied", number)
	return max
}
