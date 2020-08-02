package main

// 16ms (min 8ms)
// 5.9MB
func findMagicIndex(nums []int) int {
	for i := 0; i < len(nums); {
		num := nums[i]
		if i == num {
			return i
		}

		if i < num {
			i = num
			continue
		}
		i++
	}
	return -1
}

// divide and conquer
// the array is just sorted, not monotonically incremental
// divide and conquer can't be better than just transversal
// 16ms
// 5.9MB
func findMagicIndexDAC(nums []int) int {
	return findDAC(nums, 0, len(nums))
}

// [left, right)
func findDAC(nums []int, left, right int) int {
	if left < nums[left] {
		left = nums[left]
	}
	if left >= right {
		return -1
	}

	mid := (left + right) / 2
	magicL := findDAC(nums, left, mid)
	if magicL >= 0 {
		return magicL
	}
	if mid == nums[mid] {
		return mid
	}
	return findDAC(nums, mid+1, right)
}
