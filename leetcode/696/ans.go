package main

// 8ms
// 5.9MB
func countBinarySubstrings(s string) int {
	length := len(s)
	i := 0
	var count, lastCnt, curCnt int
	for i < length {
		lastCnt = curCnt
		curCnt = 1
		ch := s[i]
		i++
		for i < length && s[i] == ch {
			curCnt++
			i++
		}
		count += min(lastCnt, curCnt)
	}
	return count
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
