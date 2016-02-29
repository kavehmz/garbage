package main

import (
	"fmt"
	"regexp"
	"runtime/debug"
)

func main() {
	defer func() {
		if r := recover(); r != nil {
			// This is stupid. I am just play with it to see how it works.
			s := string(debug.Stack())
			r := regexp.MustCompile(`.*\.(go:*)`)
			fmt.Println(r.ReplaceAllString(s, "$1"))
		}
	}()

	a := 1
	b := 0
	fmt.Println(a / b)
}
