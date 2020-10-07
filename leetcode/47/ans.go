package main

import "sort"

// 0ms
// 3.4MB
func permuteUnique(nums []int) [][]int {
	size := len(nums)
	sort.Ints(nums)
	used := make([]bool, size)
	ans := make([][]int, 0)
	cand := make([]int, size)
	var dfs func(int)
	dfs = func(i int) {
		for j := 0; j < size; j++ {
			if used[j] || (j > 0 && !used[j-1] && nums[j] == nums[j-1]) {
				continue
			}
			cand[i] = nums[j]
			if i == size-1 {
				a := make([]int, size)
				copy(a, cand)
				ans = append(ans, a)
			} else {
				used[j] = true
				dfs(i + 1)
				used[j] = false
			}
		}
	}
	dfs(0)
	return ans
}
