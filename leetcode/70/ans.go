package main

// 0ms
// 2MB
// memo is optimal for multi calls
func climbStairsMemo(n int) int {
	if n <= 1 {
		return 1
	}

	if ans, ok := cache[n]; ok {
		return ans
	}

	ans := climbStairsMemo(n-1) + climbStairsMemo(n-2)
	cache[n] = ans
	return ans
}

var (
	cache = map[int]int{}
)

// 0ms
// 1.9MB
func climbStairs(n int) int {
	last, this := 0, 1
	for ; n > 0; n-- {
		last, this = this, last+this
	}
	return this
}

// TODO: Try matrix multiplication.
