package main

import (
	"fmt"
	"time"
)

type foodGo struct {
	taste string
	brand string
	pond  int64
}

func (food *foodGo) feedLion(amount int64) {
	if amount == 0 {
		amount = 1
	}
	food.pond = food.pond - amount
}

func main() {

	t := time.Now().UnixNano()
	for i := 0; i < 100000; i++ {
		var full foodGo = foodGo{taste: "DELICIOUS", brand: "SWEET-TREATZ", pond: 10}
		full.feedLion(0)
	}
	fmt.Println("seconds it took:", float64(time.Now().UnixNano()-t)/1000000000)

}
