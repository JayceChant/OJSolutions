package main

// 最长回文串

func longestPalindrome(s string) string {
	size := len(s)
	if size == 0 {
		return ""
	}
	longest := 1
	// [St, Ed)
	longestSt := 0
	dp := make([][]bool, size)
	for i := 0; i < size; i++ {
		dp[i] = make([]bool, size+1)
		dp[i][i] = true
		dp[i][i+1] = true
	}
	for length := 2; length <= size; length++ {
		for st := 0; st <= size-length; st++ {
			dp[st][st+length] = s[st] == s[st+length-1] && dp[st+1][st+length-1]
			if longest < length && dp[st][st+length] {
				longest = length
				longestSt = st
			}
		}
	}
	if longest == 1 {
		return s[0:1]
	}
	return s[longestSt : longestSt+longest]
}

// 4 ms, faster than 94.04% (min 0ms)
// 2.2 MB(2228kB), less than 83.82% (min 2224kB)
func longestPalindrome2(s string) string {
	if len(s) == 0 {
		return ""
	}
	low := 2
	high := len(s) + 1
	ans := s[0:1]
	longest := 1
	for low < high {
		length := (low + high) / 2
		ret := checkLength(s, length)
		if ret == "" && length-longest > 1 {
			length--
			ret = checkLength(s, length)
		}
		if ret != "" {
			ans = ret
			longest = length
			low = length + 1
		} else {
			high = length
		}
	}
	return ans
}

func checkLength(s string, length int) string {
	for start := 0; start <= len(s)-length; start++ {
		if checkPalindrome(s, start, length) {
			return s[start : start+length]
		}
	}
	return ""
}

func checkPalindrome(s string, start int, length int) bool {
	end := start + length - 1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}
	return true
}
