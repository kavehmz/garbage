package main

import (
	"fmt"
	"unicode/utf8"

	"golang.org/x/exp/utf8string"
)

const sample2 = "\xbd\xb2\x3d\xbc\x20\xe2\x8c\x98"
const sample = "کاوه"

func main() {
	fmt.Println(utf8.ValidString(sample2))
	x := fmt.Sprintf("%q", sample2)
	fmt.Println(x, utf8.ValidString(x))
	utf8string.New(sample2)

}
