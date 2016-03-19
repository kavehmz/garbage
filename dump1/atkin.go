package atkin

import (
	"math"
	"runtime"
)

func atkin4x2y2(solution []bool, x0, x1, y0, y1, m, max int, next chan bool) {
	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y += 2 {
			n := 4*x*x + y*y
			if n > max {
				continue
			}
			r := n % 60
			if r == 1 || r == 13 || r == 17 || r == 29 || r == 37 || r == 41 || r == 49 || r == 53 {
				solution[n] = !solution[n]
			}
		}
	}
	<-next

}

func atkin3x2y2(solution []bool, x0, x1, y0, y1, m, max int, next chan bool) {
	for x := x0; x <= x1; x += 2 {
		for y := y0; y <= y1; y += 2 {
			n := 3*x*x + y*y
			if n > max {
				continue
			}
			r := n % 60
			if r == 7 || r == 19 || r == 31 || r == 43 {
				solution[n] = !solution[n]
			}
		}
	}
	<-next

}

func atkin3x2ny2(solution []bool, x0, x1, y0, y1, m, max int, next chan bool) {
	for x := x0; x <= x1; x++ {
		for y := y0; y <= y1; y++ {
			if x < y {
				continue
			}
			n := 3*x*x - y*y
			if n > max {
				continue
			}
			r := n % 60
			if r == 11 || r == 23 || r == 47 || r == 59 {
				solution[n] = !solution[n]
			}
		}
	}

	<-next

}

func Atkin(max int) []int {
	var solution = make([]bool, max+1)
	m := int(math.Sqrt(float64(max))) + 1
	var ps []int = []int{2, 3, 5}

	cores := runtime.NumCPU()
	runtime.GOMAXPROCS(cores)
	next := make(chan bool, cores)

	go atkin4x2y2(solution, 1, m/2, 1, m, m, max, next)
	go atkin3x2y2(solution, 1, m, 2, m, m, max, next)
	go atkin3x2ny2(solution, 1, m, 1, m, m, max, next)

	next <- true
	next <- true
	next <- true

	for i := 0; i < cores; i++ {
		next <- true
	}

	for n := 1; n <= m; n++ {
		if solution[n] {
			for i := 1; i <= max/(n*n); i++ {
				solution[n*n*i] = false
			}
		}
	}

	for i := 2; i <= max; i++ {
		if solution[i] == true {
			ps = append(ps, i)
		}
	}

	return ps

}
