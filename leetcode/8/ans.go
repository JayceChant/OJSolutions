package main

import "math"

func myAtoi(str string) int {
	var ans int32
	var numSt, neg bool
	for _, ch := range str {
		if !numSt {
			if ch == ' ' {
				continue
			}
			if ch == '+' {
				numSt = true
				continue
			}
			if ch == '-' {
				numSt = true
				neg = true
				continue
			}
			numSt = true
		}
		if ch < '0' || ch > '9' {
			if neg {
				return int(-ans)
			}
			return int(ans)
		}
		if ans > math.MaxInt32/10 ||
			((ans == math.MaxInt32/10) &&
				((!neg && ch > '7') || (neg && ch > '8'))) {
			if neg {
				return math.MinInt32
			}
			return math.MaxInt32
		}
		ans = ans*10 + (ch - '0')
	}
	if neg {
		return int(-ans)
	}
	return int(ans)
}
