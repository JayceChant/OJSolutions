package main

// brute force O(N^2)
// O(N) (2N)
// 8ms, 96.30%
// 5MB, 9.26%
func jump2N(nums []int) int {
	length := len(nums)
	if length <= 1 {
		return 0
	}
	dp := make([]int, length)

	from := 0
	to := 1
	for to < length {
		end := from + nums[from]
		if end < to {
			from++
			continue
		}

		if end >= length {
			end = length - 1
		}

		step := dp[from] + 1
		for to <= end {
			dp[to] = step
			to++
		}
		from++
	}

	return dp[length-1]
}

// O(N)
// 8ms, 96.30%
// 4.2MB, 100%
func jump(nums []int) int {
	length := len(nums)

	step := 0
	// the range [, thisRange) within step
	thisRange := 1
	// the range [, nextRange) within step+1
	nextRange := 1
	i := 0
	for thisRange < length {
		step++
		for i < thisRange {
			farthest := i + nums[i] + 1
			if nextRange < farthest {
				nextRange = farthest
			}
			i++
		}
		thisRange = nextRange
	}
	return step
}
