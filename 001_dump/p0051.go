package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/kavehmz/prime"
)

var p []uint64 = prime.Primes(1000000)

type PrimeDigit struct {
	count int
	min   uint64
}

func getMatchedDigit(v uint64, s string) uint64 {
	var matched uint64 = 10
	var digit uint64
	base := v
	for i := len(s) - 1; i >= 0; i-- {
		digit = v % 10
		v = v / 10
		if s[i] == '1' {
			if matched == 10 {
				matched = digit
			} else {
				if matched != digit {
					return 0
				}
			}
		}
	}
	pn, _ := strconv.ParseUint(s, 10, 64)
	base = base - pn*matched
	return base
}

func checkPattern(s string) uint64 {
	m := make(map[uint64]PrimeDigit)
	for _, v := range p {
		if float64(v) < math.Pow(10, float64(len(s)-1)) {
			continue
		}
		base := getMatchedDigit(v, s)
		if base != 0 {

			if val, ok := m[base]; ok {
				val.count++
				m[base] = val
				if val.count == 8 {
					return val.min
				}
			} else {
				m[base] = PrimeDigit{count: 1, min: v}
			}

		}
	}
	return 0
}

func main() {

	for i := 1; i < 1024; i++ {
		d := checkPattern(strconv.FormatUint(uint64(i), 2))
		_ = "breakpoint"
		if d != 0 {
			fmt.Println(d)
			break
		}
	}
}
