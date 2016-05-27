package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var stdin *bufio.Reader

func init() {
	stdin = bufio.NewReader(os.Stdin)
}
func readStr(delim byte) string {
	l, _ := stdin.ReadString(delim)
	return strings.TrimSpace(l)
}
func readInt(delim byte) int {
	l, _ := stdin.ReadString(delim)
	i, _ := strconv.Atoi(strings.TrimSpace(l))
	return i
}

func permuteSlice(a []int, i int, f func([]int) bool) bool {
	if i == len(a) {
		return f(a)
	}
	for j := i; j < len(a); j++ {
		var t []int
		t = append(t, a[:i]...)
		t = append(t, a[j])
		t = append(t, a[i:j]...)
		t = append(t, a[j+1:]...)

		if permuteSlice(t, i+1, f) {
			return true
		}
	}
	return false
}

func main() {
	var a []int
	a = append(a, readInt(' '))
	a = append(a, readInt(' '))
	a = append(a, readInt(' '))
	a = append(a, readInt('\n'))

	sort.Sort(sort.Reverse(sort.IntSlice(a)))

	permuteSlice(a, 0, func(x []int) bool {
		if x[0] > 2 || (x[0] == 2 && x[1] > 3) || x[2] > 5 {
			return false
		}

		fmt.Println("Found it!:", x)
		return true
	})
}
