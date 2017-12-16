package main

import "io"

type T0 struct{}

func (T0) v0()  {}
func (*T0) p0() {}

type T1 struct{}

func (T1) v1()  {}
func (*T1) p1() {}

type T2 interface {
	v2()
	p2()
}

type T3 struct {
	T0
	*T1
	T2
}

type F struct {
	io.Reader
}

func main() {
	c := F.Read
	c(F{}, nil)
	var a T0
	b := T0.v0
	b(a)
}
