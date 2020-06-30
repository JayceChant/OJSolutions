package main

// 0ms, 100%
// 2.3MB, 100%
func longestValidParentheses(s string) int {
	var longest, matchLen, stackSize int
	for _, ch := range s {
		if ch == '(' {
			stackSize++
		} else if stackSize > 0 {
			stackSize--
			matchLen += 2
			if stackSize == 0 && matchLen > longest {
				longest = matchLen
			}
		} else { // stackSize == 0
			matchLen = 0
		}
	}

	if stackSize == 0 {
		return longest
	}

	matchLen = 0
	stackSize = 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ')' {
			stackSize++
		} else if stackSize > 0 {
			stackSize--
			matchLen += 2
			if stackSize == 0 && matchLen > longest {
				longest = matchLen
			}
		} else { // stackSize == 0
			matchLen = 0
		}
	}

	return longest
}

func longestValidParenthesesDP(s string) int {
	var longest, stackSize int
	dp := make([]int, len(s))
	for i, ch := range s {
		if ch == '(' {
			stackSize++
		} else if stackSize > 0 {
			stackSize--
			matchLen := dp[i-1] + 2
			if i > matchLen {
				matchLen += dp[i-matchLen]
			}
			if matchLen > longest {
				longest = matchLen
			}
			dp[i] = matchLen
		}
	}
	return longest
}
