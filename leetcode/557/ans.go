package main

import "unsafe"

// 4ms
// 6MB
func reverseWords(s string) string {
	slen := len(s)
	buf := make([]byte, 0, slen+1)
	start := 0
	var end int
	for start < slen {
		end = start + 1
		for end < slen && s[end] != ' ' {
			end++
		}

		for i := end - 1; i >= start; i-- {
			buf = append(buf, s[i])
		}
		buf = append(buf, ' ')
		start = end + 1
	}
	buf = buf[:slen]
	return *((*string)(unsafe.Pointer(&buf)))
}
