package main

// 1 <= nums.length <= 100
// 1 <= nums[i] <= 100
func numIdenticalPairs(nums []int) int {
	size := len(nums)
	hash := make([]int, 101)
	hash[nums[0]]++

	count := 0
	for i := 1; i < size; i++ {
		count += hash[nums[i]]
		hash[nums[i]]++
	}
	return count
}
