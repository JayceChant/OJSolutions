package builtins

import "container/list"

type entry struct {
	key   int
	value int
}

// LRUCache ...
// 144ms, 42.95% (min 116ms)
// 17.6MB, 45.45% (min 14528)
type LRUCache struct {
	queue list.List
	elems map[int]*list.Element
	cap   int
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	return LRUCache{
		elems: make(map[int]*list.Element),
		cap:   capacity,
	}
}

// Get ...
func (cache *LRUCache) Get(key int) int {
	if elem, exist := cache.elems[key]; exist {
		// move back
		cache.queue.Remove(elem)
		cache.elems[key] = cache.queue.PushBack(elem.Value)
		return elem.Value.(*entry).value
	}
	return -1
}

// Put ...
func (cache *LRUCache) Put(key int, value int) {
	if elem, exist := cache.elems[key]; exist {
		// mova back and update value
		cache.queue.Remove(elem)
		pair := elem.Value.(*entry)
		pair.value = value
		cache.elems[key] = cache.queue.PushBack(pair)
		return
	}

	if len(cache.elems) >= cache.cap {
		// eliminate
		elem := cache.queue.Front()
		cache.queue.Remove(elem)
		delete(cache.elems, elem.Value.(*entry).key)
	}
	cache.elems[key] = cache.queue.PushBack(&entry{
		key:   key,
		value: value,
	})
}
