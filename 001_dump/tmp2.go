package main

import (
	"fmt"
	"time"

	"github.com/kavehmz/prime"
)

func main() {
	s := time.Now().UnixNano()

	s = time.Now().UnixNano()
	fmt.Println(len(prime.Primes(1000000000)))
	s = time.Now().UnixNano() - s
	fmt.Println(float64(s) / 1000000.0)
}
