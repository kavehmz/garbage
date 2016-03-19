package main

import (
	"fmt"
	"regexp"
	"strconv"
	"time"
)

func main() {
	s := "text, MESSAGE 55.55"

	var m string
	var f float64
	t := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		fmt.Sscanf(s, "text, %s %f", &m, &f)
	}
	fmt.Println(m, f)
	fmt.Println("Sscanf took:", float64(time.Now().UnixNano()-t)/1000000000)

	r := regexp.MustCompile(`text, (\w+) (\d+\.\d+)`)
	t = time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		match := r.FindAllStringSubmatch(s, -1)[0]
		m = match[1]
		f, _ = strconv.ParseFloat(match[2], 64)
	}
	fmt.Println(m, f)
	fmt.Println("regexp took:", float64(time.Now().UnixNano()-t)/1000000000)

}
