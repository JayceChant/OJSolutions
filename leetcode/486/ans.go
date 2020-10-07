package main

// PredictTheWinnerMemo 0ms, 2.4MB
func PredictTheWinnerMemo(nums []int) bool {
	size := len(nums)
	cache := map[int]int{}
	var memo func(int, int) int

	// [left, right]
	memo = func(left, right int) int {
		if left == right {
			return nums[left]
		}

		hash := left*100 + right
		if res, ok := cache[hash]; ok {
			return res
		}

		res := max(nums[left]-memo(left+1, right), nums[right]-memo(left, right-1))
		cache[hash] = res
		return res
	}

	return memo(0, size-1) >= 0
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// PredictTheWinner 0ms, 2.1MB
func PredictTheWinner(nums []int) bool {
	size := len(nums)
	dp := make([]int, size)
	copy(dp, nums)
	last := make([]int, size)
	for span := 1; span < size; span++ {
		dp, last = last, dp
		k := size - span
		for left := 0; left < k; left++ {
			dp[left] = max(nums[left]-last[left+1], nums[left+span]-last[left])
		}
	}
	return dp[0] >= 0
}
