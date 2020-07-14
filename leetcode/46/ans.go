package main

// 0ms, 100%
// 2.6MB, 100%
func permute(nums []int) [][]int {
	ans := make([][]int, 0)
	size := len(nums)
	cand := make([]int, size)
	used := make([]bool, size)
	return permuteRec(nums, size, 0, used, cand, ans)
}

func permuteRec(nums []int, size int, index int, used []bool, cand []int, ans [][]int) [][]int {
	if index == size {
		a := make([]int, size)
		copy(a, cand)
		return append(ans, a)
	}
	for i := 0; i < size; i++ {
		if used[i] {
			continue
		}
		used[i] = true
		cand[index] = nums[i]
		ans = permuteRec(nums, size, index+1, used, cand, ans)
		used[i] = false
	}
	return ans
}
