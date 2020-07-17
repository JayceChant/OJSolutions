package main

// with binary search, better performce with large data
// 4ms, 69.47%
func nextPermutationBS(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	i := n - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}

	if i == 0 {
		j := n - 1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
		return
	}

	target := i - 1
	j := n
	// find min num greater than nums[target],
	// use binary search in sorted range
	for i < j-1 { // [i, j)
		m := (i + j) / 2
		if nums[m] > nums[target] {
			i = m
		} else {
			j = m
		}
	}

	nums[target], nums[i] = nums[i], nums[target]
	i = target + 1
	j = n - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}

// 0ms, 100%
func nextPermutation(nums []int) {
	n := len(nums)
	if n < 2 {
		return
	}

	i := n - 1
	for i > 0 && nums[i-1] >= nums[i] {
		i--
	}

	if i == 0 {
		j := n - 1
		for i < j {
			nums[i], nums[j] = nums[j], nums[i]
			i++
			j--
		}
		return
	}

	target := i - 1
	// find min num greater than nums[target]
	i = n - 1
	for nums[i] <= nums[target] {
		i--
	}

	nums[target], nums[i] = nums[i], nums[target]
	i = target + 1
	j := n - 1
	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
	}
}
