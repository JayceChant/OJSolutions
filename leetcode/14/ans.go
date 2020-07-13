package main

import "math"

// 0ms, 100%
func longestCommonPrefix(strs []string) string {
	n := len(strs)
	if n == 0 {
		return ""
	}

	minLen := math.MaxInt32
	for i := n - 1; i >= 0; i-- {
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
		}
	}
	pos := 0
posLoop:
	for pos < minLen {
		for i := n - 1; i >= 1; i-- {
			if strs[i][pos] != strs[i-1][pos] {
				break posLoop
			}
		}
		pos++
	}
	return strs[0][:pos]
}
