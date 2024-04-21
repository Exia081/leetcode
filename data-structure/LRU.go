package main

import (
	"container/list"
	"fmt"
)

type LRUCache struct {
	size             int
	doubleLinkedList *list.List
	cacheMap         map[interface{}]*list.Element
}

type entry struct {
	key   interface{}
	value interface{}
}

func NewLRUCache(size int) *LRUCache {
	return &LRUCache{
		size:             size,
		doubleLinkedList: list.New(),
		cacheMap:         make(map[interface{}]*list.Element, size),
	}
}

func (c *LRUCache) Get(key interface{}) (value interface{}, ok bool) {
	if ent, ok := c.cacheMap[key]; ok {
		c.doubleLinkedList.MoveToFront(ent)
		return ent.Value.(*entry).value, true
	}
	return
}

func (c *LRUCache) Add(key, value interface{}) {
	if ent, ok := c.cacheMap[key]; ok {
		c.doubleLinkedList.MoveToFront(ent)
		ent.Value.(*entry).value = value
		return
	}

	ent := &entry{key, value}
	entry := c.doubleLinkedList.PushFront(ent)
	c.cacheMap[key] = entry

	if c.doubleLinkedList.Len() > c.size {
		c.removeOldest()
	}
}

func (c *LRUCache) removeOldest() {
	ent := c.doubleLinkedList.Back()
	if ent != nil {
		c.removeElement(ent)
	}
}

func (c *LRUCache) removeElement(e *list.Element) {
	c.doubleLinkedList.Remove(e)
	kv := e.Value.(*entry)
	delete(c.cacheMap, kv.key)
}

func main() {
	cache := NewLRUCache(2)
	cache.Add("1", 1)
	cache.Add("2", 2)
	fmt.Println(cache.Get("1")) // will return 1
	cache.Add("3", 3)           // this operation will evict key 2
	fmt.Println(cache.Get("2")) // will return 0 (not found)
	cache.Add("4", 4)           // this operation will evict key 1
	fmt.Println(cache.Get("1")) // will return 0 (not found)
	fmt.Println(cache.Get("3")) // will return 3
	fmt.Println(cache.Get("4")) // will return 4
}
