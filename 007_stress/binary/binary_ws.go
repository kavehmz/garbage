package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

// var req = `{
// "proposal": 1,
// "amount": "100",
// "basis": "payout",
// "contract_type": "CALL",
// "currency": "USD",
// "duration": "60",
// "duration_unit": "s",
// "symbol": "R_100"
// }`

var req = `
{
	"website_status": 1
}`

func chk(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	for i := 0; i < 2000; i++ {
		go func(num int) {
			t := time.Now().UnixNano()
			ws, err := websocket.Dial("wss://ws.binaryws.com/websockets/v3?app_id=1&l=EN", "", "http://load_test")
			chk(err)
			fmt.Println("Time", num, ":", float64(time.Now().UnixNano()-t)/1000000000)

			for j := 0; j < 3000; j++ {
				t = time.Now().UnixNano()
				_, err = ws.Write([]byte(req))
				chk(err)
				var msg = make([]byte, 5120)
				_, err := ws.Read(msg)
				chk(err)
				// fmt.Println(string(msg))
				if j%10 == 0 {
					ws, _ = websocket.Dial("wss://ws.binaryws.com/websockets/v3?app_id=1&l=EN", "", "http://load_test")
					fmt.Printf("[%d, %d] Received for : %d after %f\n", j, num, num, float64(time.Now().UnixNano()-t)/1000000000)
				}
				time.Sleep(time.Millisecond * 1100)
			}
		}(i)
	}

	time.Sleep(time.Second * 10000)
}
