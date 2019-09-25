package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

var connections int

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
	r, _ := url.Parse("http://localhost:4003")

	t0 := time.Now()
	done := make(chan bool)
	go func() {
		defer func() {
			r := recover()
			fmt.Println("request to backend failed:", r)
		}()

		connections++
		httputil.NewSingleHostReverseProxy(r).ServeHTTP(res, req)
		connections--
		done <- true
	}()

	select {
	case <-done:
	case <-time.After(time.Second):
		log.Println("Timeout", connections, time.Since(t0), req.URL.String())
	}
}

func main() {
	log.Println("Starting the proxy server")
	http.HandleFunc("/", handleRequestAndRedirect)
	if err := http.ListenAndServe(":4002", nil); err != nil {
		panic(err)
	}
}
