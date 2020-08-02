package main

// 0ms
// 2.3MB
func subsets(nums []int) (ans [][]int) {
	size := len(nums)
	full := 1<<size - 1
	for mask := 0; mask <= full; mask++ {
		set := make([]int, 0)
		m := mask
		for i := 0; m > 0; i++ {
			if m&1 == 1 {
				set = append(set, nums[i])
			}
			m >>= 1
		}
		ans = append(ans, set)
	}
	return
}
