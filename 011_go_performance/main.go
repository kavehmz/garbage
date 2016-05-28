package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"time"
)

type p struct {
	X, Y, Z int
	Name    string
}

func goVSJSON() {
	b := bytes.Buffer{}
	enc := gob.NewEncoder(&b)
	dec := gob.NewDecoder(&b)
	t := time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		enc.Encode(p{i, 4, 5, "Test"})
		var v p
		dec.Decode(&v)
	}
	fmt.Println("Enc/Dec gob time (ns/op) :", float64(time.Now().UnixNano()-t)/1000000)

	t = time.Now().UnixNano()
	for i := 0; i < 1000000; i++ {
		bb, _ := json.Marshal(p{i, 4, 5, "Test"})
		var v p
		json.Unmarshal(bb, &v)
	}
	fmt.Println("Enc/Dec json time (ns/op) :", float64(time.Now().UnixNano()-t)/1000000)
}

func main() {
	goVSJSON()
}
