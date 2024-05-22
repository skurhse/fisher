package lrucache

import (
	"container/list"
)

type LRUCache struct {
	capacity int
	store    map[int]*list.Element
	usage    *list.List
}

type LRUEntry struct {
	key   int
	value int
}

func New(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		store:    make(map[int]*list.Element, capacity),
		usage:    list.New(),
	}
}

func (cache *LRUCache) Get(key int) int {
	element, ok := cache.store[key]

	if ok {
		cache.usage.MoveToFront(element)
		return element.Value.(*LRUEntry).value
	}

	return -1
}

func (cache *LRUCache) Put(key int, value int) {
	element, ok := cache.store[key]

	if ok {
		cache.usage.MoveToFront(element)
		element.Value.(*LRUEntry).value = value
		return
	}

	if cache.usage.Len() == cache.capacity {
		backElement := cache.usage.Back()
		cache.usage.Remove(backElement)

		backKey := backElement.Value.(*LRUEntry).key
		delete(cache.store, backKey)

	}

	entry := &LRUEntry{key: key, value: value}

	element = cache.usage.PushFront(entry)

	cache.store[key] = element
	return
}
