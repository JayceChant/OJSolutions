package main

import "strings"

// 8ms
// 5MB
func repeatedSubstringPatternBf(s string) bool {
	slen := len(s)
	plen := 1
	pi := 0
	ni := 1
	for plen <= slen/2 && ni < slen {
		pi = 0
		for pi < plen && ni < slen {
			if s[pi] != s[ni] {
				plen++
				for plen <= slen/2 && slen%plen != 0 {
					plen++
				}
				ni = plen
				break
			}
			pi++
			ni++
		}
	}
	return pi == plen
}

// 12ms
// 6.5MB
func repeatedSubstringPatternStr(s string) bool {
	slen := len(s)
	for plen := 1; plen <= slen/2; plen++ {
		if slen%plen == 0 && s[plen:]+s[:plen] == s {
			return true
		}
	}
	return false
}

// 0ms
// 6MB
func repeatedSubstringPattern(s string) bool {
	return strings.Count(s[1:]+s[:len(s)-1], s) > 0
}
