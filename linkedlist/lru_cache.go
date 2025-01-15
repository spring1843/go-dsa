package linkedlist

import (
	"container/list"
	"fmt"
	"strings"
)

type (
	lruCache struct {
		list     *list.List
		elements map[int]*list.Element
		cap      int
	}
	element struct {
		key   int
		value int
	}
)

func newLRU(capacity int) *lruCache {
	return &lruCache{
		list:     list.New(),
		elements: make(map[int]*list.Element),
		cap:      capacity,
	}
}

// get solves the problem in O(1) time and O(n) space.
func (cache *lruCache) get(key int) int {
	if _, ok := cache.elements[key]; ok {
		v := cache.elements[key]
		cache.list.MoveToBack(v)
		return v.Value.(*element).value
	}
	return -1
}

// put solves the problem in O(1) time and O(n) space.
func (cache *lruCache) put(key int, value int) {
	if v, ok := cache.elements[key]; ok {
		v.Value.(*element).value = value
		cache.list.MoveToBack(v)
		return
	}

	cache.list.PushBack(&element{key: key, value: value})
	cache.elements[key] = cache.list.Back()

	if len(cache.elements) > cache.cap {
		delete(cache.elements, cache.list.Front().Value.(*element).key)
		cache.list.Remove(cache.list.Front())
	}
}

// String represents the cache as a string.
func (cache *lruCache) String() string {
	var pairs []string
	for e := cache.list.Front(); e != nil; e = e.Next() {
		pairs = append(pairs, fmt.Sprintf("%d:%d", e.Value.(*element).key, e.Value.(*element).value))
	}
	return strings.Join(pairs, ",")
}
