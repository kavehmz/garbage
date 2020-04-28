package main

import "fmt"

type item struct {
	data       interface{}
	prev, next *item
}

type queue struct {
	start, end *item
}

func (m *queue) add(data interface{}) {
	item := item{data: data}
	if m.end == nil {
		m.start = &item
		m.end = &item
		return
	}

	item.prev = m.end
	m.end.next = &item
	m.end = &item
}

func (m *queue) remove() interface{} {
	if m.start == nil {
		return nil
	}
	item := m.start

	if m.start == m.end {
		m.start = nil
		m.end = nil
	} else {
		m.start = m.start.next
	}
	return item.data
}

func main() {
	q := queue{}
	q.add(10)
	q.add(11)
	q.add(12)

	fmt.Println(q.remove())
	fmt.Println(q.remove())
	fmt.Println(q.remove())
	fmt.Println(q.remove())
	q.add(14)
	fmt.Println(q.remove())
	q.add(15)
	fmt.Println(q.remove())
	fmt.Println(q.remove())
}
