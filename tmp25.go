package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/websocket"
)

func echoServer(ws *websocket.Conn) {
	request := make([]byte, 10000)
	n, err := ws.Read(request)
	if err != nil {
		fmt.Println("Error", err.Error())
		return
	}
	rStr := string(request[:n])
	fmt.Println("Read:", n, "It was ", rStr, " type:", ws.Request().FormValue("l"))
	ws.Write([]byte(rStr))
}

func hs(c *websocket.Config, r *http.Request) error {
	fmt.Println("Handshaking")
	return nil
}

// This example demonstrates a trivial echo server.
func main() {
	fmt.Println("before")
	ws := websocket.Server{}
	ws.Handshake = hs
	ws.Handler = echoServer
	http.Handle("/images/", http.StripPrefix("/images/", http.FileServer(http.Dir("./static"))))
	http.Handle("/websockets/v3", ws)
	fmt.Println("assigned")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic("ListenAndServe: " + err.Error())
	}
}
