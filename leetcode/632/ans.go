package main

import (
	"container/heap"
	"math"
)

// 56ms
// 6.7MB
func smallestRange(nums [][]int) (ans []int) {
	k := len(nums)
	h := make(minHeap, k)
	max := math.MinInt32
	for i := 0; i < k; i++ {
		num := nums[i][0]
		h[i] = []int{num, i, 0, len(nums[i])}
		if max < num {
			max = num
		}
	}
	heap.Init(&h)

	minWidth := math.MaxInt32
	ans = make([]int, 2)
	for {
		min := h[0]
		diff := max - min[0]
		if minWidth > diff {
			minWidth = diff
			ans[0] = min[0]
			ans[1] = max
		}

		min[2]++
		if min[2] == min[3] {
			break
		}

		num := nums[min[1]][min[2]]
		if max < num {
			max = num
		}
		min[0] = num
		heap.Fix(&h, 0)
	}

	return
}

var _ heap.Interface = (*minHeap)(nil)

// []int{num, arr_i, num_i, arr_size}
type minHeap [][]int

func (a minHeap) Len() int            { return len(a) }
func (a minHeap) Swap(i, j int)       { a[i], a[j] = a[j], a[i] }
func (a minHeap) Less(i, j int) bool  { return a[i][0] < a[j][0] }
func (a *minHeap) Push(x interface{}) { *a = append(*a, x.([]int)) }
func (a *minHeap) Pop() interface{} {
	old := *a
	newSize := len(old) - 1
	*a = old[:newSize]
	return old[newSize]
}
