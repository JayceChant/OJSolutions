package main

func isMatch(s string, p string) bool {
	return isMatchDP2(s, p)
}

func isMatchBT(s string, p string) bool {
	return backTracking(s, p, len(s)-1, len(p)-1)
}

func backTracking(s string, p string, is int, ip int) bool {
	for { // 普通匹配循环
		if ip < 0 {
			if is < 0 { // 两边同时结束，没有发现不匹配
				return true
			}
			// pattern 结束而 string 有剩余，不匹配
			// 反之 只有 is < 0 需要进一步判断，因为剩余的 pattern 有可能匹配空串
			return false
		}

		if p[ip] == '*' { // '*' 进入递归回溯，为避免嵌套太深，break 后另起 for 循环
			break
		}

		// 1. pattern 不是 '*'，而 string 还有剩余，不匹配
		// 2. 字符不匹配
		if is < 0 || (s[is] != p[ip] && p[ip] != '.') {
			return false
		}
		is--
		ip--
	}

	ip--
	for { // '*' 回溯循环
		if backTracking(s, p, is, ip-1) { // 剩余串匹配
			return true
		}
		// 剩余串不匹配，'*' 尝试再匹配一个字符后递归

		// 1. string 结束，'*' 无法进一步匹配
		// 2. 字符不匹配
		if is < 0 || (s[is] != p[ip] && p[ip] != '.') {
			return false
		}
		is--
	}
}

type state uint8

const (
	unknown state = iota
	unmatched
	matched
)

var memo [][]state

func isMatchMemoRL(s string, p string) bool {
	memo = make([][]state, len(s)+1)
	for i := range memo {
		memo[i] = make([]state, len(p)+1)
		memo[i][0] = unmatched
	}
	memo[0][0] = matched
	return memoRL(s, p, len(s), len(p))
}

func memoRL(s string, p string, is int, ip int) bool {
	if memo[is][ip] == unknown {
		if p[ip-1] == '*' {
			if memoRL(s, p, is, ip-2) ||
				(is >= 1 && (p[ip-2] == '.' || s[is-1] == p[ip-2]) && memoRL(s, p, is-1, ip)) {
				memo[is][ip] = matched
			} else {
				memo[is][ip] = unmatched
			}
		} else {
			if is >= 1 && (p[ip-1] == '.' || s[is-1] == p[ip-1]) &&
				memoRL(s, p, is-1, ip-1) {
				memo[is][ip] = matched
			} else {
				memo[is][ip] = unmatched
			}
		}
	}
	return memo[is][ip] == matched
}

func isMatchMemo(s string, p string) bool {
	memo = make([][]state, len(s)+1)
	for i := range memo {
		memo[i] = make([]state, len(p)+1)
		memo[i][len(p)] = unmatched
	}
	memo[len(s)][len(p)] = matched
	return memoDP(s, p, 0, 0)
}

func memoDP(s string, p string, is int, ip int) bool {
	if memo[is][ip] == unknown {
		preMatch := is < len(s) && (p[ip] == '.' || s[is] == p[ip])

		if ip+1 < len(p) && p[ip+1] == '*' {
			if memoDP(s, p, is, ip+2) ||
				(preMatch && memoDP(s, p, is+1, ip)) {
				memo[is][ip] = matched
			} else {
				memo[is][ip] = unmatched
			}
		} else {
			if preMatch &&
				memoDP(s, p, is+1, ip+1) {
				memo[is][ip] = matched
			} else {
				memo[is][ip] = unmatched
			}
		}
	}
	return memo[is][ip] == matched
}

func isMatchDP(s string, p string) bool {
	dp := make([][]bool, len(s)+1)
	for i := range dp {
		dp[i] = make([]bool, len(p)+1)
	}
	dp[len(s)][len(p)] = true

	for is := len(s); is >= 0; is-- {
		for ip := len(p) - 1; ip >= 0; ip-- {
			preMatch := is < len(s) && (p[ip] == '.' || s[is] == p[ip])
			if ip+1 < len(p) && p[ip+1] == '*' {
				dp[is][ip] = dp[is][ip+2] ||
					(preMatch && dp[is+1][ip])
			} else {
				dp[is][ip] = preMatch && dp[is+1][ip+1]
			}
		}
	}

	return dp[0][0]
}

func isMatchDP2(s string, p string) bool {
	if len(p) == 0 {
		return len(s) == 0
	}

	lastDp := make([]bool, len(p)+1)
	curDp := make([]bool, len(p)+1)
	for is := len(s); is >= 0; is-- {
		for ip := len(p) - 1; ip >= 0; ip-- {
			preMatch := is < len(s) && (p[ip] == '.' || s[is] == p[ip])
			if ip+1 < len(p) && p[ip+1] == '*' {
				curDp[ip] = curDp[ip+2] ||
					(preMatch && lastDp[ip]) || (is == len(s) && ip+2 == len(p))
			} else {
				curDp[ip] = preMatch && (lastDp[ip+1] || (is+1 == len(s) && ip+1 == len(p)))
			}
		}
		lastDp, curDp = curDp, lastDp
	}

	return lastDp[0]
}
