package main

import (
	"math/rand"
)

func findKthLargest(nums []int, k int) int {
	head := 0
	tail := len(nums)
	k-- // index starts from 0
	for {
		piv := partition(nums, head, tail)
		if piv == k {
			return nums[k]
		} else if piv > k {
			tail = piv
		} else { // piv < k
			head = piv + 1
		}
	}
}

// [head, tail)
func partition(nums []int, head int, tail int) int {
	if head+1 == tail {
		return head
	}
	x := rand.Intn(tail - head)
	nums[head], nums[head+x] = nums[head+x], nums[head]

	piv := head
	head++
	tail--
	for {
		for head <= tail && nums[head] >= nums[piv] {
			head++
		}
		for nums[tail] < nums[piv] {
			tail--
		}
		if head >= tail {
			break
		}
		nums[head], nums[tail] = nums[tail], nums[head]
	}
	nums[piv], nums[tail] = nums[tail], nums[piv]

	return tail
}

type heap []int

func (h heap) build() {
	n := len(h)
	for parent := n/2 - 1; parent >= 0; parent-- {
		h.down(parent)
	}
}

func (hp *heap) pop() int {
	old := *hp
	n := len(old)
	top := old[0]
	old[0] = old[n-1]
	// bubble down
	old.down(0)
	*hp = old[:n-1]
	return top
}

func (h heap) down(i int) {
	n := len(h)
	for {
		child := i*2 + 1
		if child >= n {
			break
		}
		if child+1 < n && h[child] < h[child+1] {
			child++
		}
		if h[i] >= h[child] {
			break
		}
		h[i], h[child] = h[child], h[i]
		i = child
	}
}

// 4ms, 99,42%
// 3.5MB, 100%
func findKthLargestHeap(nums []int, k int) int {
	h := heap(nums)
	h.build()
	for i := 1; i < k; i++ {
		h.pop()
	}
	return h.pop()
}
