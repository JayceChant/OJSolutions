package main

import "container/heap"

// 68ms (min 44ms)
// 6.8MB
func isPossible(target []int) bool {
	if len(target) == 1 {
		return target[0] == 1
	}

	h := intHeap(target)
	heap.Init(&h)
	var sum, restSum int
	for i := len(h) - 1; i >= 1; i-- {
		restSum += h[i]
	}

	for h[0] > 1 {
		if h[0]-restSum < 1 {
			return false
		}
		h[0] %= restSum
		if restSum != 1 && h[0] == 0 {
			return false
		}
		sum = h[0] + restSum

		heap.Fix(&h, 0)
		restSum = sum - h[0]
	}
	return true
}

var _ heap.Interface = (*intHeap)(nil)

type intHeap []int

func (a intHeap) Len() int            { return len(a) }
func (a intHeap) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a intHeap) Less(i, j int) bool  { return a[i] > a[j] } // maxHeap
func (a *intHeap) Push(x interface{}) { *a = append(*a, x.(int)) }
func (a *intHeap) Pop() interface{} {
	old := *a
	newSize := len(old) - 1
	*a = old[:newSize]
	return old[newSize]
}
