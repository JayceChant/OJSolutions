package main

import "sort"

// 4ms, 99.32%
// 4.1MB, 100%
func searchRange(nums []int, target int) []int {
	size := len(nums)
	if size == 0 {
		return []int{-1, -1}
	}
	// [low, high]
	low := 0
	high := size - 1
	// find leftmost
	for low < high {
		mid := (low + high) / 2
		if nums[mid] == target {
			high = mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}

	if nums[low] != target {
		return []int{-1, -1}
	}

	ans := []int{low, 0}

	// [low, high)
	high = size
	for low < high-1 {
		mid := (low + high) / 2
		if nums[mid] == target {
			low = mid
		} else if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid
		}
	}
	ans[1] = low
	return ans
}

// 8ms, 92.51%
// 4.1MB, 100%
func searchRangeLib(nums []int, target int) []int {
	size := len(nums)
	if size == 0 {
		return []int{-1, -1}
	}

	low := sort.Search(size, func(i int) bool {
		return nums[i] >= target
	})
	if low >= size || nums[low] != target {
		return []int{-1, -1}
	}

	high := low + sort.Search(size-low, func(i int) bool {
		return nums[low+i] > target
	}) - 1

	return []int{low, high}
}
