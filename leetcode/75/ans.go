package main

// 0ms
// 2.1MB
func sortColors(nums []int) {
	size := len(nums)

	r0 := 0
	for r0 < size && nums[r0] == 0 {
		r0++
	}
	if r0 == size {
		return
	}

	l2 := size - 1
	for l2 >= 0 && nums[l2] == 2 {
		l2--
	}
	if l2 < 0 {
		return
	}

	i := r0 + 1
	for i <= l2 {
		switch nums[i] {
		case 0:
			nums[i] = nums[r0]
			nums[r0] = 0
			r0++
			for nums[r0] == 0 {
				r0++
			}
			if i <= r0 {
				i = r0 + 1
			}
		case 1:
			i++
		case 2:
			nums[i] = nums[l2]
			nums[l2] = 2
			l2--
			for nums[l2] == 2 {
				l2--
			}
		}
	}
	if r0 < l2 && nums[r0] == 2 {
		nums[r0], nums[l2] = nums[l2], nums[r0]
	}
}
