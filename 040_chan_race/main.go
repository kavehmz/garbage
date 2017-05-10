package main

import "fmt"
import "time"

var c = make(chan bool)

func a() {
	<-c
	fmt.Println("A")
}

func b() {
	<-c
	fmt.Println("B")

}
func main() {
	go a()
	go b()

	time.Sleep(time.Second)
	c <- true
	c <- true
	time.Sleep(time.Second)

}
