package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

var req = `
{
  "symbol": "frxAUDUSD",
  "subscribe": 1,
  "duration": 2,
  "currency": "USD",
  "amount": 10,
  "proposal": 1,
  "duration_unit": "m",
  "basis": "payout",
  "contract_type": "PUT"
}`

func chk(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func main() {
	for i := 0; i < 235; i++ {
		go func(num int) {
			t := time.Now().UnixNano()
			ws, err := websocket.Dial("wss://www2.binary.com/websockets/v3?l=EN&debug=1", "", "http://localhost/")
			chk(err)
			fmt.Println("Time", num, ":", float64(time.Now().UnixNano()-t)/1000000000)

			for j := 0; j < 10; j++ {
				t = time.Now().UnixNano()
				_, err = ws.Write([]byte(req))
				chk(err)
				var msg = make([]byte, 5120)
				_, err = ws.Read(msg)
				chk(err)
				fmt.Printf("Received: %d after %f\n", num, float64(time.Now().UnixNano()-t)/1000000000)
			}
		}(i)
	}

	time.Sleep(time.Second * 50)
}
