package main

// 8ms
// 4.2MB
func canJump(nums []int) bool {
	target := len(nums) - 1
	if target <= 0 {
		return true
	}
	maxEnd := 0
	for i := 0; i <= maxEnd; i++ {
		end := i + nums[i]
		if maxEnd < end {
			maxEnd = end
			if maxEnd >= target {
				return true
			}
		}
	}
	return false
}
