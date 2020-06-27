package main

// 4 ms, 93.16%
// 3.4 MB, 61.77%
func twoSum(nums []int, target int) []int {
	numToIdx := make(map[int]int, len(nums))
	ans := make([]int, 0, 2)
	for i, num := range nums {
		i2, ok := numToIdx[target-num]
		if ok {
			ans = append(ans, i2, i)
			break
		}
		numToIdx[num] = i
	}
	return ans
}
