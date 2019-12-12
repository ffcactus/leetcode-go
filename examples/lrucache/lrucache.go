package lrucache

import (
	"container/list"
	"fmt"
)

type cacheElement struct {
	k int
	v int
}

type LRUCache struct {
	capacity int
	tree     map[int]*list.Element
	list     *list.List
}

func Constructor(capacity int) LRUCache {
	fmt.Printf("cache := Constructor(%d)\n", capacity)
	return LRUCache{
		capacity: capacity,
		tree:     make(map[int]*list.Element, capacity),
		list:     list.New(),
	}
}

func (this *LRUCache) Get(key int) int {
	fmt.Printf("cache.Get(%d)\n", key)
	le, ok := this.tree[key]
	if !ok {
		return -1
	}
	this.list.MoveToFront(le)
	return le.Value.(*cacheElement).v
}

func (this *LRUCache) Put(key int, value int) {
	fmt.Printf("cache.Put(%d, %d)\n", key, value)
	oldElement, ok := this.tree[key]
	if ok {
		e := oldElement.Value.(*cacheElement)
		e.v = value
		this.list.MoveToFront(oldElement)
		return
	}

	if this.list.Len() == this.capacity {
		elementToDelete := this.list.Back()
		// delete from tree
		delete(this.tree, elementToDelete.Value.(*cacheElement).k)
		// delete from list
		this.list.Remove(elementToDelete)
	}

	listElement := this.list.PushFront(&cacheElement{
		v: value,
		k: key,
	})
	this.tree[key] = listElement
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
