package main

import (
	"context"
	"fmt"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	r := rate.NewLimiter(1, 5)

	for i := 0; i < 10; i++ {
		fmt.Println(i)
		if r.AllowN(time.Now(), 1) {
			fmt.Println("ggg", i)
		}
		fmt.Println("wait")
		r.WaitN(context.Background(), 1)
		fmt.Println("wait done")

		res := r.ReserveN(time.Now(), 1)
		if !res.OK() {
			// Not allowed to act! Did you remember to set lim.burst to be > 0 ?
			return
		}
		time.Sleep(res.Delay())
	}
}
