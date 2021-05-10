package main

import (
	"container/heap"
	"log"
)

type ItemData []int

type Item struct {
	value    ItemData
	distance int
	index    int
}

type PQType []*Item

func (p PQType) Len() int {
	return len(p)
}

func (p PQType) Less(i, j int) bool {
	return p[i].distance > p[j].distance
}

func (p PQType) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = j
	p[j].index = i
}

func (p *PQType) Push(x interface{}) {
	n := len(*p)
	item := x.(*Item)
	item.index = n
	*p = append(*p, item)
}

func (p *PQType) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*p = old[0 : n-1]
	return item
}

func closeness(c []int) int {
	return c[0]*c[0] + c[1]*c[1]
}

func kClosest(points [][]int, k int) [][]int {
	ret := [][]int{}
	if len(points) == 0 || k == 0 {
		return ret
	}

	set := &PQType{}
	heap.Init(set)

	for i := 0; i < len(points); i++ {
		newItem := &Item{
			value:    ItemData(points[i]),
			distance: closeness(points[i]),
		}
		if k > len(*set) {
			heap.Push(set, newItem)
			continue
		}

		// worst := heap.Pop(set).(*Item)
		if (*set)[0].distance > newItem.distance {
			_ = heap.Pop(set)
			heap.Push(set, newItem)
		}
	}
	n := len(*set)
	for i := 0; i < n; i++ {
		tmp := heap.Pop(set).(*Item)
		ret = append(ret, tmp.value)
	}

	return ret
}

func main() {
	// in := [][]int{{1, 5}, {8, 9}, {8, 9}}
	i := [][]int{{3, 3}, {5, -1}, {-2, 4}}
	log.Println(kClosest(i, 2))

}
