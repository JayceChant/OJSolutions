package main

import "math"

func reverse(x int) int {
	neg := (x < 0)
	if neg {
		x = -x
	}
	ans := 0
	for x > 0 {
		ans = ans*10 + (x % 10)
		x /= 10
	}
	if neg {
		ans = -ans
		if ans < math.MinInt32 {
			return 0
		}
	} else {
		if ans > math.MaxInt32 {
			return 0
		}
	}
	return ans
}

// Not exactly on point.
// If the environment can only store 32 bits, it has overflowed before judging.
func reverse2(x int) int {
	ans := 0
	for x != 0 {
		ans = ans*10 + (x % 10)
		if ans < math.MinInt32 || math.MaxInt32 < ans {
			return 0 // early prune
		}
		x /= 10
	}
	return ans
}

func reverse3(x int) int {
	ans := 0
	for x != 0 {
		if ans < math.MinInt32/10 || math.MaxInt32/10 < ans {
			return 0 // early prune
		}
		ans = ans*10 + (x % 10)
		x /= 10
	}
	return ans
}

// double-100%
// use int32 reduce 4kB
func reverse4(x int) int {
	ans := int32(0)
	for x != 0 {
		if ans < math.MinInt32/10 || math.MaxInt32/10 < ans {
			return 0 // early prune
		}
		ans = ans*10 + int32(x%10)
		x /= 10
	}
	return int(ans)
}
