package main

import "container/heap"

// 0ms
// 2MB
func lastStoneWeight(stones []int) int {
	h := maxHeap(stones)
	heap.Init(&h)

	for len(h) >= 2 {
		a := h[0]
		heap.Pop(&h)
		b := h[0]
		heap.Pop(&h)
		c := a - b
		if c > 0 {
			heap.Push(&h, c)
		}
	}

	if len(h) == 1 {
		return h[0]
	}
	// len(h) == 0
	return 0
}

type maxHeap []int

var _ heap.Interface = (*maxHeap)(nil)

func (a maxHeap) Len() int            { return len(a) }
func (a maxHeap) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a maxHeap) Less(i, j int) bool  { return a[i] > a[j] }
func (a *maxHeap) Push(x interface{}) { *a = append(*a, x.(int)) }
func (a *maxHeap) Pop() interface{} {
	old := *a
	newSize := len(old) - 1
	*a = old[:newSize]
	return old[newSize]
}
