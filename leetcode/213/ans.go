package main

// house robber II
// DP
// 0ms
// 2MB
func rob(nums []int) int {
	length := len(nums)
	if length == 0 {
		return 0
	}

	if length == 1 {
		return nums[0]
	}

	// act and skip from index 0
	act0 := nums[0]
	skip0 := 0
	// act and skip from index 1
	act1 := 0
	skip1 := 0
	for i := 1; i < length-1; i++ {
		num := nums[i]
		act0, skip0 = skip0+num, max(act0, skip0)
		act1, skip1 = skip1+num, max(act1, skip1)
	}
	return max(max(act0, skip0), skip1+nums[length-1])
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
