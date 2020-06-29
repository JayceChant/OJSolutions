package main

import "math"

func minSubArrayLen(s int, nums []int) int {
	minLen := math.MaxInt32
	sum := 0
	left := 0
	right := 0
	for {
		// 谨慎扩张(<)，激进收缩(>=)
		for right < len(nums) && sum < s {
			sum += nums[right]
			right++
		}
		if sum < s {
			break
		}
		for sum >= s {
			sum -= nums[left]
			left++
		}
		length := right - left + 1 // [left, right)
		if length == 1 {
			return 1
		}
		if length < minLen {
			minLen = length
		}
	}
	if minLen == math.MaxInt32 {
		return 0
	}
	return minLen
}
