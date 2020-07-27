package main

// 0ms, 100%
// 2MB, 100%
func isSubsequence(s string, t string) bool {
	si := 0
	for ti := 0; ti < len(t) && si < len(s); ti++ {
		if s[si] == t[ti] {
			si++
		}
	}

	if si == len(s) {
		return true
	}
	return false
}

// pre-compute for multi s
// 0ms
// 6.6MB
func isSubsequenceDP(s string, t string) bool {
	sizeT := len(t)

	if sizeT == 0 {
		return len(s) == 0
	}

	var dp [][26]int

	dp, ok := cache[t]
	if !ok {
		dp = make([][26]int, sizeT)
		for ch := 0; ch < 26; ch++ {
			dp[sizeT-1][ch] = sizeT
		}

		dp[sizeT-1][t[sizeT-1]-'a'] = sizeT - 1
		for i := sizeT - 2; i >= 0; i-- {
			for ch := 0; ch < 26; ch++ {
				dp[i][ch] = dp[i+1][ch]
			}
			dp[i][t[i]-'a'] = i
		}
		cache[t] = dp
	}

	si := 0
	ti := 0
	for ; si < len(s) && ti < len(t); si++ {
		ti = dp[ti][s[si]-'a'] + 1
	}

	if si == len(s) && ti <= sizeT {
		return true
	}
	return false
}

var (
	cache = map[string]([][26]int){}
)
