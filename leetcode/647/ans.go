package main

// 0ms
// 2MB

func countSubstrings(s string) int {
	length := len(s)
	count := length
	extLen := length*2 - 1 // add '#' at odd index
MID_LOOP:
	for mid := 1; mid < extLen-1; mid++ {
		offset := 2 - mid%2
		end := min(mid, extLen-mid-1)
		for offset <= end {
			if s[(mid-offset)/2] != s[(mid+offset)/2] {
				continue MID_LOOP
			}
			count++
			offset += 2
		}
	}
	return count
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
