package main

import . "fmt"

const (
	a = iota + iota
	b = iota + iota
)

const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Partyday
	numberOfDays // this constant is not exported
)

func main() {
	Println(numberOfDays)
}
