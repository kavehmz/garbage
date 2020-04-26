package main

import (
	"errors"
	"fmt"
	"log"
	"reflect"
)

func main() {
	x := &stack{}
	x.push(5)
	err := x.push(2.2)
	log.Println("expect 2.2 to be rejected:", err)
	x.push(10)

	fmt.Println(x.pop()) // 5
	fmt.Println(x.pop()) // 10
	fmt.Println(x.pop()) // nil
}

var defaultData interface{}

type item struct {
	previous *item
	data     interface{}
}

type stack struct {
	current *item
}

func (m *stack) push(data interface{}) error {
	if m.current != nil && reflect.TypeOf(data).Name() != reflect.TypeOf(m.current.data).Name() {
		return errors.New("types don't match")
	}
	pushItem := item{
		data:     data,
		previous: m.current,
	}
	m.current = &pushItem
	return nil

}
func (m *stack) pop() interface{} {
	if m.current == nil {
		return nil
	}
	popItem := m.current
	m.current = m.current.previous
	popItem.previous = nil
	return popItem.data
}
