package main

// 4ms, 89.66%
// 3.1MB, 100%
// relative to 33, 153
func findMin(nums []int) int {
	size := len(nums)
	low := 0
	high := size - 1
	for low < high {
		if nums[low] < nums[high] {
			return nums[low]
		}

		mid := (low + high) / 2
		if nums[mid] == nums[high] {
			if nums[low] == nums[mid] {
				low++
				high--
				continue
			}

			if nums[low] > nums[mid] {
				high = mid
			} else {
				low = mid + 1
			}
			continue
		}

		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}
