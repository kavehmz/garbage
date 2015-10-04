/*
How the answer is not 42 also!
Can be solved mathematically but that way wont help me checking golang
*/
package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

var tnm map[int]int
var MAX_LEN int = 20

func init() {
	tnm = make(map[int]int)

	for i := 0; i < MAX_LEN*26; i++ {
		tnm[i] = i * (i + 1) / 2
	}
}

func is_triangle_number(s int) bool {
	for i := 0; i < len(tnm); i++ {
		if tnm[i] == s {
			return true
		}
	}
	return false
}

func main() {
	content, _ := ioutil.ReadFile("p042_words.txt")
	r, _ := regexp.Compile("\"([^\"]+)\",?")
	for _, word := range r.FindAllStringSubmatch(string(content), -1) {
		if len(word[1]) > MAX_LEN {
			panic(fmt.Sprintf("Length of %s more than supported MAX_LEN(%d)", word[1], MAX_LEN))
		}
		score := 0
		for _, letter := range word[1] {
			score += int(letter) - 64
		}
		fmt.Printf("%s [%d]: %t\n", word[1], score, is_triangle_number(score))
	}

}
