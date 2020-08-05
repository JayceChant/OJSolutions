package main

// house robber
// DP
// 0ms
// 2MB
func rob(nums []int) int {
	act := 0
	skip := 0
	for _, num := range nums {
		act, skip = skip+num, max(act, skip)
	}
	return max(act, skip)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
