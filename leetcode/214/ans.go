package main

import "unsafe"

// 0ms
// 3.3MB (min 2852kB)
func shortestPalindrome(s string) string {
	if s == "" {
		return ""
	}

	slen := len(s)
	exlen := slen*2 - 1
	manacher := make([]int, exlen)
	expand := func(l, r int) int {
		if l&1 == 0 {
			l--
			r++
		}
		l--
		r++

		for l >= 0 && r < exlen && s[l/2] == s[r/2] {
			l -= 2
			r += 2
		}
		return (r - l - 4) / 2
	}

	longest := 1
	right := -1
	j := -1
	for i := 1; i < slen; i++ {
		var armLen int
		if right >= i {
			minArmLen := min(manacher[j*2-i], right-i)
			armLen = expand(i-minArmLen, i+minArmLen)
		} else {
			armLen = expand(i, i)
		}
		manacher[i] = armLen

		if i+armLen > right {
			j = i
			right = i + armLen
		}
		if armLen == i {
			longest = i + 1
		}
	}
	buf := make([]byte, 0, slen*2-longest)
	for i := slen - 1; i >= longest; i-- {
		buf = append(buf, s[i])
	}
	buf = append(buf, s...)
	return *(*string)(unsafe.Pointer(&buf))
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
