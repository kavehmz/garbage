package main

import "fmt"

type A struct {
	a int
	b int
}

func (a *A) pp() {
	fmt.Println("AAA")
}

type B struct {
	a string
}

type BB struct {
	B
}

type C struct {
	A
	B
	a int
}

type D struct {
	A
	C
}

func main() {
	b := BB{B{"dd"}}

	c := C{A: A{1, 2}, B: B{"1"}}
	c.pp()
	c.A.a = 3
	c.B.a = "tt"
	c.b = 4
	fmt.Println(c.b, b)
	fmt.Println(c.A.a)
	fmt.Println(c)
	var d D
	fmt.Println(d)
}
