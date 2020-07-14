package main

// no duplicate
// 0ms, 100%
// 2.6MB, 71.43%
func search(nums []int, target int) int {
	size := len(nums)
	if size == 0 {
		return -1
	}
	low := 0
	high := size - 1
	// assuming the nums[min] is min
	for low < high {
		mid := (low + high) / 2
		//fmt.Println("1:", low, mid, high)
		if target == nums[mid] {
			return mid
		}

		// nums[low] <= target < nums[mid]
		// maybe mid==low==0, do not access nums[mid-1]
		if nums[low] <= target && target < nums[mid] {
			high = mid - 1
			break
		}

		// nums[mid+1] <= target <= nums[high]
		// (mid, mid+1) may be the jump point
		if nums[mid+1] <= target && target <= nums[high] {
			low = mid + 1
			break
		}

		// nums[min] <= target <= nums[high] < nums[low]
		// or
		// nums[high] < nums[low] <= target <= nums[min-1]
		if nums[mid] > nums[high] {
			// mid < min
			low = mid + 1
		} else {
			// min < mid
			high = mid - 1
		}
	}
	for low < high {
		mid := (low + high) / 2
		//fmt.Println("2:", low, mid, high)
		if target == nums[mid] {
			return mid
		}

		if target < nums[mid] {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if target == nums[low] {
		return low
	}
	return -1
}

// func findMin(nums []int) int {
// 	size := len(nums)
// 	low := 0
// 	high := size - 1
// 	for low != high {
// 		mid := (low + high) / 2
// 		if nums[mid] > nums[high] {
// 			low = mid
// 		} else {
// 			high = mid - 1
// 		}
// 	}
// 	return low
// }
