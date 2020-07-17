package main

// 8ms, 89.89% (min 0ms)
func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	tail := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			continue
		}
		if tail < i {
			nums[tail] = nums[i]
		}
		tail++
	}
	return tail
}
