//usr/bin/env go run  -gcflags '-m -l' $0 $@;exit
package main

import "fmt"

type s struct {
	i *int32
}

func main() {
	var x s
	x.i = new(int32)
	_ = escapesToHeap(x)
	_ = wontEscapesToHeap(x)
}

func escapesToHeap(z s) *s {
	i := 42
	fmt.Println(i)
	return &z
}

func wontEscapesToHeap(y s) s {
	return y
}
