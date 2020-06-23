package main

var (
	gs     string
	gwords []string
	wcount int
	wlen   int
	wused  []bool
)

func findSubstringBT(s string, words []string) []int {
	gs = s
	gwords = words
	slen := len(gs)
	wcount = len(gwords)
	ans := []int{}
	if slen == 0 || wcount == 0 {
		return ans
	}

	wlen = len(gwords[0])
	subtotal := wlen * wcount
	wused = make([]bool, wcount)
	for start := 0; start+subtotal <= slen; start++ {
		if dfs(start, 0) {
			ans = append(ans, start)
		}
	}

	return ans
}

func dfs(start, wi int) bool {
	si := start + wlen*wi
	for wj := 0; wj < wcount; wj++ {
		if wused[wj] {
			continue
		}

		if gs[si:si+wlen] != gwords[wj] {
			continue
		}

		if wi == wcount-1 {
			// last word matched
			return true
		}

		// forwards
		wused[wj] = true
		// 只要有一个匹配，剩下的没有必要再试；
		// 再试除了浪费时间，还会导致答案出现重复。
		matched := dfs(start, wi+1)
		wused[wj] = false
		if matched {
			return true
		}
	}
	return false
}

var dp [][]bool

func findSubstringDP(s string, words []string) []int {
	gs = s
	slen := len(gs)
	wcount = len(words)
	ans := []int{}
	if slen == 0 || wcount == 0 {
		return ans
	}

	wlen = len(words[0])
	subtotal := wlen * wcount
	wused = make([]bool, wcount)
	dp = make([][]bool, slen-wlen+1)
	for i := range dp {
		dp[i] = make([]bool, wcount)
	}

	for start := 0; start+wlen <= len(gs); start++ {
		for wi := 0; wi < wcount; wi++ {
			text := gs[start : start+wlen]
			if text == words[wi] {
				dp[start][wi] = true
			}
		}
	}

	for start := 0; start+subtotal <= slen; start++ {
		if dfsDP(start, 0) {
			ans = append(ans, start)
		}
	}
	return ans
}

func dfsDP(start, wi int) bool {
	for wj := 0; wj < wcount; wj++ {
		if wused[wj] {
			continue
		}

		if !dp[start][wj] {
			continue
		}

		if wi == wcount-1 {
			return true
		}

		wused[wj] = true
		matched := dfsDP(start+wlen, wi+1)
		wused[wj] = false
		if matched {
			return true
		}
	}
	return false
}

type state uint8

const (
	stateUnknown state = iota
	stateUnmatched
	stateMatched
)

var memo [][]state

func findSubstringMemo(s string, words []string) []int {
	gs = s
	gwords = words
	slen := len(gs)
	wcount = len(gwords)
	ans := []int{}
	if slen == 0 || wcount == 0 {
		return ans
	}

	wlen = len(words[0])
	subtotal := wlen * wcount
	wused = make([]bool, wcount)
	memo = make([][]state, slen-wlen+1)
	for i := range memo {
		memo[i] = make([]state, wcount)
	}

	for start := 0; start+subtotal <= slen; start++ {
		if dfsMemo(start, 0) {
			ans = append(ans, start)
		}
	}
	return ans
}

func dfsMemo(start, wi int) bool {
	for wj := 0; wj < wcount; wj++ {
		if wused[wj] {
			continue
		}

		if memo[start][wj] == stateUnknown {
			if gs[start:start+wlen] == gwords[wj] {
				memo[start][wj] = stateMatched
			} else {
				memo[start][wj] = stateUnmatched
			}
		}

		if memo[start][wj] != stateMatched {
			continue
		}

		if wi == wcount-1 {
			return true
		}

		wused[wj] = true
		matched := dfsMemo(start+wlen, wi+1)
		wused[wj] = false
		if matched {
			return true
		}
	}
	return false
}

// 子串计算器，因为不确定子串数，uint8 只到 255，用 uint16
var word2count map[string]uint16

func findSubstringHash(s string, words []string) []int {
	gs = s
	wcount := len(words)
	ans := make([]int, 0, 32)
	if len(gs) == 0 || wcount == 0 {
		return ans
	}

	wlen = len(words[0])
	subtotal := wlen * wcount

	word2count = make(map[string]uint16)
	for _, word := range words {
		word2count[word]++
	}

	for start := 0; start+subtotal <= len(gs); start++ {
		if dfsHash(start, wcount) {
			ans = append(ans, start)
		}
	}
	return ans
}

func dfsHash(start, wordLeft int) bool {
	word := gs[start : start+wlen]
	count, ok := word2count[word]
	if !ok || count <= 0 {
		return false
	}

	if wordLeft == 1 {
		return true
	}

	word2count[word]--
	matched := dfsHash(start+wlen, wordLeft-1)
	word2count[word]++
	return matched
}
