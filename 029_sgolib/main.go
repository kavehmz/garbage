package main

import (
	"errors"
	"fmt"
	"log"
)

// MyLib is a template to show one type of error handling in a lib.
type MyLib struct {
	Error func(error)
	i     int
}

func (s *MyLib) die() {
	s.checkError(errors.New("I will die now"))
}

func (s *MyLib) checkError(e error) {
	if e == nil {
		return
	}
	if s.Error != nil {
		s.Error(e)
		return
	}
	log.Fatal(e)
}

func main() {
	var m MyLib
	m.Error = func(e error) {
		fmt.Println(e)
	}

	m.die()
}
