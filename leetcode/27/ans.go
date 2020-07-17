package main

// 0ms
func removeElement(nums []int, val int) int {
	count := 0
	for i := range nums {
		if nums[i] == val {
			continue
		}
		if count < i {
			nums[count] = nums[i]
		}
		count++
	}
	return count
}
