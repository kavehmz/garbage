package problems

import (
	"sync"
	"time"
)

// Element is general
type Element interface{}

type timeSet struct {
	members map[Element]time.Time
	sync.RWMutex
}

func (s *timeSet) init() {
	s.Lock()
	defer s.Unlock()
	s.members = make(map[Element]time.Time)
}

func (s *timeSet) len() int {
	s.RLock()
	defer s.RUnlock()
	return len(s.members)
}

func (s *timeSet) get(e Element) (time.Time, bool) {
	s.RLock()
	defer s.RUnlock()
	val, ok := s.members[e]
	return val, ok
}

func (s *timeSet) set(e Element, t time.Time) {
	s.Lock()
	s.members[e] = t
	s.Unlock()
}

func (s *timeSet) list() []Element {
	s.RLock()
	defer s.RUnlock()
	l := make([]Element, 0, s.len())
	for k := range s.members {
		l = append(l, k)
	}
	return l
}

// WWL type is useful
type WWL struct {
	add    timeSet
	remove timeSet
	sync.RWMutex
}

// Init sets
func (wwl *WWL) Init() *WWL {
	wwl.add.init()
	wwl.remove.init()
	return wwl
}

// Add sets
func (wwl *WWL) Add(e Element, t time.Time) {
	if val, ok := wwl.add.get(e); !ok || t.UnixNano() > val.UnixNano() {
		wwl.add.set(e, t)
	}
}

// Remove sets
func (wwl *WWL) Remove(e Element, t time.Time) {
	if val, ok := wwl.remove.get(e); !ok || t.UnixNano() > val.UnixNano() {
		wwl.remove.set(e, t)
	}
}

// Exists sets
func (wwl *WWL) Exists(e Element) bool {
	a, aok := wwl.add.get(e)
	r, rok := wwl.remove.get(e)
	if !rok {
		return aok
	}
	return a.UnixNano() > r.UnixNano()
}

// Get sets
func (wwl *WWL) Get() []Element {

	var l []Element
	for _, e := range wwl.add.list() {
		if wwl.Exists(e) {
			l = append(l, e)
		}
	}
	return l
}
