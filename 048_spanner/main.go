package main

import (
	"flag"
	"fmt"
	"sort"
	"sync"
	"time"

	"gonum.org/v1/gonum/floats"
	"gonum.org/v1/gonum/stat"
)

var cun *int
var max *int
var settle *int

func main() {
	max = flag.Int("max", 0, "data number int")
	cun = flag.Int("cun", 0, "concurrency int")
	settle = flag.Int("settle", 0, "in ms. settle time before,for larger concurrencies which init might affect the stats")
	flag.Parse()

	if spannerProject != "" {
		benchmark(spannerInsert, spannerSelect, spannerSelect)
	}

	if redisURL != "" {
		benchmark(redisInsert, nil, nil)
	}

	if aerospikeAddr != "" {
		benchmark(aerospikeInsert, aerospikeSelect, nil)
	}
}

func benchmark(in, selnew, selold func(int)) {

	operationStart := time.Now()
	fmt.Println("CUN", cun)
	min := time.Now().Nanosecond()

	lock := sync.Mutex{}
	w := make(chan bool, *cun)
	var s0 []float64
	for i := min; i < min+10; i++ {
		in(i)
	}

	for i := min + 10; i < min+*max+10; i++ {
		w <- true
		go func(n int) {
			if time.Since(operationStart) < time.Millisecond*time.Duration(*settle) {
				time.Sleep(time.Millisecond * time.Duration(*settle))
			}
			start := time.Now()
			in(n)
			lock.Lock()
			s0 = append(s0, float64(time.Since(start).Nanoseconds()/1000000))
			lock.Unlock()
			<-w
		}(i)
	}
	for i := 0; i < *cun; i++ {
		w <- true
	}
	stats("INSERT", s0)

	if selnew != nil {
		for i := min; i < min+10; i++ {
			selnew(i)
		}

		w1 := make(chan bool, *cun)
		var s1 []float64
		for i := min; i < min+*max; i++ {
			w1 <- true
			go func(n int) {
				if time.Since(operationStart) < time.Millisecond*time.Duration(*settle) {
					time.Sleep(time.Millisecond * time.Duration(*settle))
				}
				start := time.Now()
				selnew(n)
				lock.Lock()
				s1 = append(s1, float64(time.Since(start).Nanoseconds()/1000000))
				lock.Unlock()
				<-w1
			}(i)
		}
		for i := 0; i < *cun; i++ {
			w1 <- true
		}
		stats("SELECT NEW DATA", s1)
	}

	if selold != nil {
		w2 := make(chan bool, *cun)
		var s2 []float64
		for i := min; i < min+*max; i++ {
			w2 <- true
			go func(n int) {
				if time.Since(operationStart) < time.Millisecond*time.Duration(*settle) {
					time.Sleep(time.Millisecond * time.Duration(*settle))
				}
				start := time.Now()
				selold(min)
				lock.Lock()
				s2 = append(s2, float64(time.Since(start).Nanoseconds()/1000000))
				lock.Unlock()
				<-w2
			}(i)
		}
		for i := 0; i < *cun; i++ {
			w2 <- true
		}
		stats("SELECT OLD DATA", s2)
	}
}

func stats(title string, x []float64) {
	sort.Float64s(x)
	mean := stat.Mean(x, nil)
	stdev := stat.StdDev(x, nil)
	sum := floats.Sum(x)
	stdErr := stat.StdErr(stdev, sum)
	q1 := stat.Quantile(0.01, stat.Empirical, x, nil)
	q95 := stat.Quantile(0.95, stat.Empirical, x, nil)
	q99 := stat.Quantile(0.99, stat.Empirical, x, nil)

	fmt.Println("                         ", title)
	fmt.Println("Num", len(x))
	fmt.Println("Sum", sum)
	fmt.Println("Mean", mean)
	fmt.Println("Min", x[0])
	fmt.Println("Max", x[len(x)-1])
	fmt.Println("StdDev", stdev)
	fmt.Println("stdErr", stdErr)
	fmt.Println("Quantile 1pct", q1)
	fmt.Println("Quantile 95pct", q95)
	fmt.Println("Quantile 99pct", q99)
}
