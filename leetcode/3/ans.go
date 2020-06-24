package main

// if existed
// 8 ms, faster than 75.98%
// 3.2 MB, less than 48.30%
// if idx > start
// 4 ms, faster than 92.89%
// 3.3 MB, less than 43.09%
func lengthOfLongestSubstringMap(s string) int {
	var longest, start int
	charToIdx := make(map[rune]int)
	for i, char := range s {
		idx := charToIdx[char]
		if idx > start {
			length := i - start
			if length > longest {
				longest = length
			}
			start = idx
		}
		charToIdx[char] = i + 1
	}
	length := len(s) - start
	if length > longest {
		longest = length
	}
	return longest
}

// 4 ms, faster than 92.89%
// 2.6 MB, less than 90.94%
func lengthOfLongestSubstring(s string) int {
	var longest, start int
	chars := []byte(s)
	// assume that s contains ascii only
	charToIdx := make([]int, 128)
	for i, char := range chars {
		idx := charToIdx[char]
		if idx > start {
			length := i - start
			if length > longest {
				longest = length
			}
			start = idx
		}
		// to distinguish from default 0, record index + 1
		charToIdx[char] = i + 1
	}
	length := len(s) - start
	if length > longest {
		longest = length
	}
	return longest
}
