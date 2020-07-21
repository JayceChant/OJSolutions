package main

// 0ms, 100%
// 2.6MB, 14%
// relative to 33, 154
func findMin(nums []int) int {
	size := len(nums)
	low := 0
	high := size - 1
	for low != high {
		mid := (low + high) / 2
		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}
