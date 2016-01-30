package main

import "fmt"

type ByteSize int64

const (
	_           = iota // ignore first value by assigning to blank identifier
	KB ByteSize = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
)

func main() {
	fmt.Println(EB)
}
