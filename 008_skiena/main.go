package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kavehmz/garbage/008_skiena/problems"
)

type a struct {
	name   string
	family string
}

var w problems.WWL

func aa(a interface{}) {
	fmt.Print(a)
}

func main() {
	var r int
	r = 9223372036854775807
	fmt.Println(r)
	fmt.Println(time.Now().UnixNano())
	w := problems.WWL{}
	fmt.Print(w)
	fmt.Println()
	aa(w)
	w.Init()
	e := a{name: "some"}
	w.Add(e, time.Now().UTC())
	w.Add(e, time.Now().UTC())
	fmt.Println(w.Exists(e))
	e = a{name: "test"}
	p := (w.Get()[0]).(a)
	fmt.Println(e, p)
}

func bst() {
	a := make([]int, 0, 400000)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < cap(a); i++ {
		a = append(a, r.Int())
	}

	var root *problems.Node
	for i := 0; i < len(a); i++ {
		problems.Insert(&root, a[i])
	}

	problems.Iterate(root, func(node *problems.Node) {
		fmt.Println(node.Value)
	})

	problems.BSTCount = 0
	problems.Find(root, a[cap(a)-1])
	fmt.Println(problems.BSTCount)
}

func junk01() {
	fmt.Println(problems.Divide(49, 7))

	a := []int{1, 4, 3, 1, 7, 9, 2, 8}
	problems.Merge(a)
	fmt.Println("Merge", a)

	fmt.Println(problems.Fib(32))
	fmt.Println(problems.Steps)

	fmt.Println(problems.Fib(32))

	a = []int{1, 4, 3, 1, 7, 9, 2, 8}
	problems.Quick(a)
	fmt.Println("Quick", a)

	a = []int{1, 4, 3, 1, 7, 9, 2, 8}
	problems.Heap(a)
	fmt.Println("Heap", a)
}
