package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"plugin"
	"sync"
	"time"

	"github.com/golang/glog"
)

var myApp struct {
	app *plugin.Plugin
	sync.RWMutex
	handler *func(http.ResponseWriter, *http.Request)
}

func handler(w http.ResponseWriter, r *http.Request) {
	if myApp.handler == nil {
		fmt.Fprintf(w, "Hi there, my app is not ready yet")
	}
	(*myApp.handler)(w, r)
}

func loadMyApp() {
	timer := time.NewTicker(time.Second)
	for {
		<-timer.C
		glog.Info("Checking time ", time.Now().Unix())
		fname := fmt.Sprintf("%d.so", time.Now().Unix())
		if _, err := os.Stat(fname); os.IsNotExist(err) {
			continue
		}
		p, err := plugin.Open(fname)
		if err != nil {
			log.Println(err)
			continue
		}
		glog.Info("Plugin loaded: ", fname)

		f, err := p.Lookup("MyAppHandler")
		if err != nil {
			log.Println(err)
			continue
		}
		glog.Info("Func loaded")

		myApp.app = p
		myApp.Lock()
		h := f.(func(http.ResponseWriter, *http.Request))
		myApp.handler = &h
		myApp.Unlock()
	}
}

func main() {
	flag.Parse()
	go loadMyApp()
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
