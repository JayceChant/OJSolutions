package main

// 0ms
// 2.1MB
func wiggleMaxLength(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return length
	}

	l := 0
	r := 1
	for r < length && nums[l] == nums[r] {
		r++
	}
	if r == length {
		return 1
	}

	subLen := 2 // at least [l, r]
	lt1 := nums[l] < nums[r]
	l = r
	r++
	for r < length {
		if nums[l] == nums[r] {
			r++
			continue
		}

		// nums[l] != nums[r]
		lt2 := nums[l] < nums[r]
		if lt1 != lt2 {
			subLen++
			l = r
			lt1 = lt2
		} else {
			l++
		}
		r++
	}

	return subLen
}

// fail case
// [1,17,5,10,13,15,10,16,8]
