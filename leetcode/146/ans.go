package main

// LRUCache ...
// 128ms, 96.44% (min 116ms)
// 14.4MB, 100%
type LRUCache struct {
	entries map[int]*node
	queue   *deque
	cap     int
}

// Constructor ...
func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		entries: make(map[int]*node),
		queue:   new(deque),
		cap:     capacity,
	}
	cache.queue.init()
	return cache
}

// Get ...
func (cache *LRUCache) Get(key int) int {
	if entry, exist := cache.entries[key]; exist {
		cache.queue.remove(entry)
		cache.queue.push(entry)
		return entry.value
	}
	return -1
}

// Put ...
func (cache *LRUCache) Put(key int, value int) {
	if entry, exist := cache.entries[key]; exist {
		cache.queue.remove(entry)
		entry.value = value
		cache.queue.push(entry)
		return
	}

	var entry *node
	if cache.queue.size == cache.cap {
		// reuse the eliminated node
		entry = cache.queue.pop()
		delete(cache.entries, entry.key)
		entry.key = key
		entry.value = value
	} else {
		entry = &node{
			key:   key,
			value: value,
		}
	}
	cache.queue.push(entry)
	cache.entries[key] = entry
}

type node struct {
	key   int
	value int
	prev  *node
	next  *node
}

type deque struct {
	// sentinel root
	// root.next == head
	// root.prev == tail
	root node
	size int
}

func (q *deque) init() {
	sentinel := &(q.root)
	q.root.next = sentinel
	q.root.prev = sentinel
}

func (q *deque) push(e *node) {
	e.prev = q.root.prev
	e.prev.next = e
	e.next = &(q.root)
	q.root.prev = e
	q.size++
}

func (q *deque) remove(e *node) {
	if q.size == 0 {
		return
	}
	e.prev.next = e.next
	e.next.prev = e.prev
	q.size--
	// avoid memory leaks
	e.prev = nil
	e.next = nil
}

func (q *deque) pop() *node {
	e := q.root.next // head
	q.remove(e)
	return e
}
