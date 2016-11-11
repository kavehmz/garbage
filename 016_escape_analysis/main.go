//usr/bin/env go run  -gcflags '-m -l' $0 $@;exit
package main

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
	return &z
}

func wontEscapesToHeap(y s) s {
	return y
}
