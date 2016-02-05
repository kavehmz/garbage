package main

import (
	"fmt"
	"io"
)
import "github.com/kavehmz/enoox"

type setupMessage struct {
	ready   bool
	message string
}

// func setupUiServer(dbReady chan bool) {
// 	ins := createDigitalOcteanInstance(myUISetup)
// 	dbState := <-dbReady
// 	if dbState {
// 		bootInstace(ins)
// 	}
// }
//
// func setupClientDB(dbReady chan string) {
// 	ins := createDigitalOcteanInstance(myDBSetup)
// 	bootInstace(ins)
// 	initDBData(ins)
// 	if checkDBState() {
// 		dbReady <- true
// 		return
// 	}
// 	dbReady <- false
// 	return
// }

func aa(aa enoox.Server) {
	io.Writer
	return
}

func main() {
	fmt.Println()
	var a enoox.RectAA
	fmt.Println(a)
	// com := make(chan string, 0)
	// go setupUiServer(com)
	// go setupClientDB(com)
}
