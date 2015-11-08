package main

import (
	"fmt"
	"time"

	//"./atkin"

	"github.com/kavehmz/prime"
)

func main() {

	for k, v := range []string{"id", "account_id", "action_type", "referrer_id", "contract_id", "payment_id", "amount", "balance_after"} {
		fmt.Println(k, v)
	}

	max := 1000000
	s := time.Now().UnixNano()

	fmt.Println(17 / 2)

	//	var m runtime.MemStats
	//
	//	runtime.ReadMemStats(&m)
	//	fmt.Println("Mem: ", m)

	s = time.Now().UnixNano()
	fmt.Println(len(prime.Primes(uint64(max))))
	s = time.Now().UnixNano() - s
	fmt.Println(float64(s) / 1000000.0)

	//	runtime.ReadMemStats(&m)
	//	fmt.Println("Mem: ", m.Alloc)

	//	s = time.Now().UnixNano()
	//	fmt.Println(len(atkin.Atkin(max)))
	//	s = time.Now().UnixNano() - s
	//	fmt.Println(float64(s) / 1000000.0)
	//
	//	runtime.ReadMemStats(&m)
	//	fmt.Println("Mem: ", m.Alloc)

}
