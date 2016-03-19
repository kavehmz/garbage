package main

import (
	"fmt"
	"math/rand"
	"time"
)

type state func(*person) state

type person struct {
	hungry int
	tired  int
	curios int
	age    int
}

func next(p *person) state {
	_ = "breakpoint"
	if p.age >= 100 {
		return nil
	}
	if p.hungry > 10 {
		return eat
	}
	if p.tired > 10 {
		return sleep
	}
	if p.curios > 10 {
		return study
	}
	return eat
}

func eat(p *person) state {
	fmt.Println("eat")
	p.hungry -= rand.Intn(p.hungry)
	p.tired++
	p.curios++
	p.age++

	return next(p)
}

func sleep(p *person) state {
	fmt.Println("sleep")
	p.hungry++
	p.tired = 0
	p.curios++
	p.age++

	return next(p)
}

func study(p *person) state {
	fmt.Println("study")
	p.hungry++
	p.tired++
	p.curios -= rand.Intn(p.curios)
	p.age++

	return next(p)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	p := person{hungry: 10}
	_ = "breakpoint"
	for s := next(&p); s != nil; s = s(&p) {
	}

	fmt.Printf("When the person ended he was : %+v \n", p)
}
