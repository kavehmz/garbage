package main

import (
	"container/heap"
	"sort"
)

type MeetingInterval struct {
	begin int
	end   int
}

type Room struct {
	value    MeetingInterval
	priority int
	index    int
}

type PQType []*Room

func (p PQType) Len() int {
	return len(p)
}

func (p PQType) Less(i, j int) bool {
	return p[i].priority < p[j].priority
}

func (p PQType) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
	p[i].index = j
	p[j].index = i
}

func (p *PQType) Push(x interface{}) {
	n := len(*p)
	item := x.(*Room)
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

func minMeetingRooms(meetings [][]int) int {
	if len(meetings) == 0 {
		return 0
	}

	sort.Slice(meetings, func(i, j int) bool {
		if meetings[i][0] == meetings[j][0] {
			return meetings[i][1] > meetings[j][1]
		}
		return meetings[i][0] < meetings[j][0]
	})

	scheduledRooms := &PQType{}
	heap.Init(scheduledRooms)
	heap.Push(scheduledRooms, &Room{
		value:    MeetingInterval{meetings[0][0], meetings[0][1]},
		priority: meetings[0][1],
		index:    0})

	for i := 1; i < len(meetings); i++ {
		nextMetting := meetings[i]
		earliestRoom := heap.Pop(scheduledRooms).(*Room)

		if earliestRoom.value.end <= nextMetting[0] {
			earliestRoom.value = MeetingInterval{nextMetting[0], nextMetting[1]}
			earliestRoom.priority = nextMetting[1]
			heap.Push(scheduledRooms, earliestRoom)
			continue
		} else {
			heap.Push(scheduledRooms, earliestRoom)
		}

		heap.Push(scheduledRooms, &Room{
			value:    MeetingInterval{nextMetting[0], nextMetting[1]},
			priority: nextMetting[1],
		})
	}
	return len(*scheduledRooms)
}

func main() {
	// in := [][]int{{1, 5}, {8, 9}, {8, 9}}
	// in := [][]int{{0, 30}, {5, 10}, {15, 20}}
	// in := [][]int{{9, 10}, {4, 9}, {4, 17}}
	in := [][]int{{2, 11}, {6, 16}, {11, 16}}

}
