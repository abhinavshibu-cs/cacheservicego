package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	capacity int
	cache    map[int]*list.Element
	list     *list.List
}

type Pair struct {
	key   int
	value string
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[int]*list.Element),
		list:     list.New(),
	}
}

func (l *LRUCache) Get(key int) string {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(Pair).value
	}
	return ""
}

func (l *LRUCache) Put(key int, value string) {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		elem.Value = Pair{key, value}
	} else {
		if l.list.Len() >= l.capacity {
			evicted := l.list.Back()
			l.list.Remove(evicted)
			delete(l.cache, evicted.Value.(Pair).key)
		}
		l.cache[key] = l.list.PushFront(Pair{key, value})
	}
}

func main() {
	cache := NewLRUCache(3)
	cache.Put(1, "A")
	cache.Put(2, "B")
	cache.Put(3, "C")
	fmt.Println(cache.Get(1))
	cache.Put(4, "D")
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
}
