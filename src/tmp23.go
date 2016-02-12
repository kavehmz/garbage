package main

import "fmt"

type t struct {
	m string
}

func (a t) aa() {
	fmt.Println("test")
}

func main() {
	t{m: "test"}.aa()
}
