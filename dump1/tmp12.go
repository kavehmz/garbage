package main

import (
	"fmt"
	"runtime"
)

func changeSample(p **int) {
	*p = new(int)
	**p = 8
}

func main() {
	var i *int
	fmt.Println(4 * *2)
	fmt.Println(i)
	changeSample(&i)
	runtime.Breakpoint()
	fmt.Println(*i)
}

//https://play.golang.org/p/Vhk5D6XYqP
//https://play.golang.org/p/Cm3cQ-0PjX
