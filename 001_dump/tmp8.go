package main

import "fmt"

type AA struct {
	a int
}

func sum(i interface{}) {
	fmt.Println(i.(int))
}

func main() {
	var i int = 2
	sum(i)
}
