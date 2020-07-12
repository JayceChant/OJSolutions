package main

import (
	"container/heap"
	"math"
	"sort"
)

var _ heap.Interface = (*maxHeap)(nil)

func minDifference(nums []int) int {
	n := len(nums)
	if n <= 4 {
		return 0
	}

	// !! min use maxHeap, max use minHeap, to leave required elements to the last
	min4 := make(maxHeap, 4)
	max4 := make(minHeap, 4)
	copy(min4, nums)
	copy(max4, nums)
	heap.Init(&min4)
	heap.Init(&max4)
	for i := 4; i < n; i++ {
		heap.Push(&min4, nums[i])
		heap.Pop(&min4)
		heap.Push(&max4, nums[i])
		heap.Pop(&max4)
	}
	sort.Ints(min4)
	sort.Ints(max4)
	minDiff := math.MaxInt32
	for i := 0; i < 4; i++ {
		diff := max4[i] - min4[i]
		if diff < minDiff {
			minDiff = diff
		}
	}
	return minDiff
}

type minHeap []int

func (h minHeap) Len() int { return len(h) }

func (h minHeap) Less(i, j int) bool { return h[i] < h[j] }

func (h minHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *minHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *minHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

type maxHeap []int

func (h maxHeap) Len() int { return len(h) }

func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }

func (h maxHeap) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *maxHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *maxHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
