package main

// variation of the matrix multiplication problem

// 4ms, 94.74%
// 2.9MB, 100%
func maxCoins(nums []int) int {
	size := len(nums)
	if size == 0 {
		return 0
	}

	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	size += 2

	dp := make([][]int, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]int, size)
	}

	for step := 2; step < size; step++ {
		left := 0
		right := left + step
		for right < size {
			max := 0
			for mid := left + 1; mid < right; mid++ {
				coins := dp[left][mid] + dp[mid][right] + nums[left]*nums[mid]*nums[right]
				if max < coins {
					max = coins
				}
			}
			dp[left][right] = max
			left++
			right++
		}
	}
	//fmt.Println(dp)
	return dp[0][size-1]
}
