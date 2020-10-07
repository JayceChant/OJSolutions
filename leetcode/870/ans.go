package main

import "sort"

// 84ms
// 6.6MB
func advantageCount(A []int, B []int) []int {
	size := len(A)
	bidx := make([]int, size)
	for i := 0; i < size; i++ {
		bidx[i] = i
	}
	sort.Ints(A)
	sort.Slice(bidx, func(i, j int) bool {
		return B[bidx[i]] < B[bidx[j]]
	})

	adv := 0
	dis := size - 1
	ans := make([]int, size)
	for _, a := range A {
		if a > B[bidx[adv]] {
			ans[bidx[adv]] = a
			adv++
		} else {
			ans[bidx[dis]] = a
			dis--
		}
	}
	return ans
}
