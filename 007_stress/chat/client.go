package main

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

func main() {
	for i := 0; i < 150; i++ {
		go func() {
			origin := "http://localhost/"
			url := "ws://localhost:8080/chat"
			ws, err := websocket.Dial(url, "", origin)
			if err != nil {
				log.Fatal(err)
			}
			if _, errs := ws.Write([]byte(`{"message":"test"}`)); err != nil {
				log.Fatal(errs)
			}
			var msg = make([]byte, 512)
			var n int
			if n, err = ws.Read(msg); err != nil {
				log.Fatal(err)
			}
			time.Sleep(time.Second * 1)
			fmt.Printf("Received: %s.\n", msg[:n])
		}()
	}

	time.Sleep(time.Second * 2)
}
