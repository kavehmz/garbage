package main

import "fmt"
import (
	"encoding/binary"
	"github.com/kavehmz/prime"
	"os"
	"time"
)

func PrimesF() []uint64 {
	var t []uint64

	f, err := os.Open("primes.dat")
	panic(err)
	defer f.Close()
	c := true
	for c {
		b := make([]byte, 4)
		_, err := f.Read(b)
		if err == nil {
			n, _ := binary.Uvarint(b)
			t = append(t, n)
		} else {
			c = false
		}
	}
	return t
}

func tmp() {
	var p []uint64
	f, _ := os.Create("/tmp/dat")
	fmt.Println("start writing")

	for _, n := range p {
		b := make([]byte, 4)
		binary.PutUvarint(b, n)
		f.Write(b)
	}
	f.Close()

	f, _ = os.Open("/tmp/dat")
	defer f.Close()
	fmt.Println("start reading")

	c := true
	for c {
		b := make([]byte, 4)
		_, err := f.Read(b)
		if err == nil {
			n, _ := binary.Uvarint(b)
			n++
			//fmt.Println(n)
		} else {
			c = false
		}

	}
	fmt.Println("end")
}
func main() {
	s := time.Now().UnixNano()
	p := prime.Primes(10)
	s = time.Now().UnixNano() - s
	fmt.Println(float64(s) / 1000000.0)
	fmt.Println(len(p))

}
