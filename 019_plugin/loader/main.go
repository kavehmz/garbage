package main

import (
	"log"
	"plugin"
)

func main() {

	p, err := plugin.Open("plugin.so")
	if err != nil {
		log.Panic(err)
	}

	f, err := p.Lookup("Say")
	if err != nil {
		log.Panic(err)
	}
	f.(func())() // prints "Hello or Bye based on which plugin is available in plugin.so"

}
