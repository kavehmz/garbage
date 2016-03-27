package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

type optFunc func(int) int

func (f optFunc) do(r int) int {
	return f(r)
}

func timing() decorator {
	return func(c operation) operation {
		return optFunc(func(n int) int {
			defer func(start time.Time) {
				log.Println("Action took:", time.Since(start).Nanoseconds())
			}(time.Now())
			return c.do(n)
		})
	}
}

func logging(l *log.Logger) decorator {
	return func(c operation) operation {
		a := optFunc(func(n int) int {
			log.Println("Cacluation sqaure of:", n)
			return c.do(n)
		})
		return a
	}
}

type operation interface {
	do(int) int
}

type decorator func(operation) operation

// wrapping the opt (operator) inside other fucntions that will do monitoring, logging or benchmarking.
func decorate(opt operation, ds ...decorator) operation {
	decorated := opt
	for _, decorate := range ds {
		decorated = decorate(decorated)
	}
	return decorated
}

func square(n int) int {
	return n * n
}

func main() {
	// All above complexity is to make this call simple!
	// You can see now we can call or operator (square) and wrapper around it
	// all other monitoring and benchmarking functions in an elegant way
	s := decorate(optFunc(square),
		logging(log.New(os.Stdout, "client: ", log.LstdFlags)),
		timing(),
	).do(8)

	fmt.Println("square of 8 is:", s)
}
