package main

// 0 ms, faster than 100.00%
// 2.2 MB, less than 83.92% (min=2176kb)
func firstMissingPositive(nums []int) int {
	for i := range nums {
		num := nums[i]
		for num > 0 && num <= len(nums) && nums[num-1] != num {
			nums[num-1], num = num, nums[num-1]
		}
	}
	for i := range nums {
		if nums[i] != i+1 {
			return i + 1
		}
	}
	return len(nums) + 1
}
