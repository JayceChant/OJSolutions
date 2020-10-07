package main

import (
	"regexp"
	"sort"
	"strconv"
)

// 28ms
// 6.1MB
func isPowerOfThreeRec(n int) bool {
	if n <= 0 {
		return false
	}

	if n == 1 {
		return true
	}

	if n%3 != 0 {
		return false
	}

	return isPowerOfThreeRec(n / 3)
}

// 112ms
// 6.8MB
func isPowerOfThreeStr(n int) bool {
	if n <= 0 {
		return false
	}
	str := strconv.FormatInt(int64(n), 3)
	match, _ := regexp.MatchString("^10*$", str)
	return match
}

var (
	hash = map[int]bool{
		1:          true,
		3:          true,
		9:          true,
		27:         true,
		81:         true,
		243:        true,
		729:        true,
		2187:       true,
		6561:       true,
		19683:      true,
		59049:      true,
		177147:     true,
		531441:     true,
		1594323:    true,
		4782969:    true,
		14348907:   true,
		43046721:   true,
		129140163:  true,
		387420489:  true,
		1162261467: true,
	}
)

// 48ms
// 6.1MB
func isPowerOfThreeHash(n int) bool {
	return hash[n]
}

const (
	maxIntPow3 int = 1162261467
)

// 32ms
// 6.1MB
func isPowerOfThreeMod(n int) bool {
	return n >= 1 && maxIntPow3%n == 0
}

var (
	cand = []int{1, 3, 9, 27, 81, 243, 729, 2187, 6561, 19683, 59049, 177147, 531441, 1594323, 4782969, 14348907, 43046721, 129140163, 387420489, 1162261467}
)

// 20ms
// 6.1MB
func isPowerOfThree(n int) bool {
	if n <= 0 {
		return false
	}
	i := sort.SearchInts(cand, n)
	return i < len(cand) && n == cand[i]
}
