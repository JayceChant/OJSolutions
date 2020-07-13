package main

import (
	"math"
	"strconv"
)

func isPalindromeItoa(x int) bool {
	if x < 0 {
		return false
	}
	s := strconv.Itoa(x)
	l := 0
	r := len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isPalindrome(x int) bool {
	if x < 0 || (x > 0 && x%10 == 0) {
		return false
	}

	re := 0
	for x > re {
		re = re*10 + x%10
		x /= 10
	}
	return x == re || x == re/10
}

// slow
func isPalindromeMath(x int) bool {
	if x < 0 {
		return false
	}

	base := int(math.Pow10(int(math.Log10(float64(x)))))
	for base > 0 {
		if x%10 != x/base {
			return false
		}
		x %= base
		x /= 10
		base /= 100
	}
	return true
}
