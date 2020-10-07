package main

import "sort"

// 0ms
// 2.1MB
func makesquare(nums []int) bool {
	size = len(nums)
	if size < 4 {
		return false
	}

	sum := 0
	for _, n := range nums {
		sum += n
	}
	if sum%4 != 0 {
		return false
	}

	edge := sum / 4
	sort.Ints(nums)
	if nums[size-1] > edge {
		return false
	}

	gn = nums
	used = make([]bool, size)
	end := size - 3
	for i := size - 1; i >= end; i-- {
		if used[i] {
			end--
			continue
		}
		used[i] = true
		if nums[i] == edge {
			continue
		}
		if !dfs(edge-nums[i], i-1) {
			return false
		}
	}
	return true
}

var (
	size int
	gn   []int
	used []bool
)

func dfs(target int, st int) bool {
	last := 0
	for i := st; i >= 0; i-- {
		if used[i] || gn[i] > target || (last != 0 && gn[i] == last) {
			continue
		}

		used[i] = true

		diff := target - gn[i]
		if diff == 0 || dfs(diff, i-1) {
			return true
		}
		last = gn[i]

		used[i] = false
	}
	return false
}
