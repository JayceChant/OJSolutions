package main

import "math"

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}

	if divisor == 1 {
		return dividend
	}

	if divisor == -1 {
		if dividend == math.MinInt32 {
			return math.MaxInt32
		}
		return -dividend
	}

	neg := false
	if dividend < 0 {
		dividend = -dividend
		neg = !neg
	}

	if divisor < 0 {
		divisor = -divisor
		neg = !neg
	}

	ans, _ := div(dividend, divisor)
	if neg {
		return -ans
	}
	return ans
}

func div(dividend int, divisor int) (int, int) {
	ans := 0
	doubleDivisor := divisor + divisor
	if dividend >= doubleDivisor {
		ans, dividend = div(dividend, doubleDivisor)
		ans += ans
	}

	for dividend >= divisor {
		dividend -= divisor
		ans++
	}
	return ans, dividend
}

// baseline
// 4ms, 61.29% (min 0ms)
// 2.5MB, 50% (min 2452)
// func divide(dividend int, divisor int) int {
//     if divisor == -1 && dividend == math.MinInt32 {
// 		return math.MaxInt32
// 	}
// 	// It's a test of system behavior, not intentional cheating.
//     return dividend / divisor
// }
