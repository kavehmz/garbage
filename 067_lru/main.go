package main

import (
	"fmt"
	"sync"
)

type item struct {
	data    interface{}
	version uint64
}

type lruCache struct {
	sync.Mutex
	cache   map[string]*item
	version uint64
	maxSize int // edge case zero
}

func new(maxSize int) *lruCache {
	return &lruCache{
		maxSize: maxSize,
		cache:   make(map[string]*item),
	}

}

func (m *lruCache) write(key string, value interface{}) {
	m.version++
	it := &item{data: value, version: m.version}
	if _, found := m.cache[key]; found || len(m.cache) < m.maxSize {
		m.cache[key] = it
		return
	}

	// find smallest version /delete add new one
	var sk string
	sv := uint64(1 << 63)
	for k, v := range m.cache {
		if v.version < sv {
			sv = v.version
			sk = k
		}
	}

	delete(m.cache, sk)
	m.cache[key] = it
}

func (m *lruCache) read(key string) interface{} {
	m.version++
	if val, found := m.cache[key]; found {
		val.version = m.version
		return val.data
	}
	return nil
}

func (m *lruCache) delete(key string) {
	delete(m.cache, key)
}

func (m *lruCache) clear() {
	m.version = 0
	m.cache = make(map[string]*item)
}

func (m *lruCache) count() int {
	return len(m.cache)
}

func testBasicCache() interface{} {
	c := new(2)
	c.write("a", 2)
	c.write("b", 5)
	return c.read("a")
}

func testNeedExpiry(key string) interface{} {
	c := new(2)
	c.write("a", 2)
	c.write("b", 5)
	c.write("c", 7)
	return c.read(key)
}

func main() {
	fmt.Println("Expecting 2", testBasicCache())
	fmt.Println("Expecting nil", testNeedExpiry("a"))
	fmt.Println("Expecting 5", testNeedExpiry("b"))
	fmt.Println("Expecting 7", testNeedExpiry("c"))
}
