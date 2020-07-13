package main

var (
	numToLetters = map[byte]([]byte){
		'2': []byte{'a', 'b', 'c'},
		'3': []byte{'d', 'e', 'f'},
		'4': []byte{'g', 'h', 'i'},
		'5': []byte{'j', 'k', 'l'},
		'6': []byte{'m', 'n', 'o'},
		'7': []byte{'p', 'q', 'r', 's'},
		'8': []byte{'t', 'u', 'v'},
		'9': []byte{'w', 'x', 'y', 'z'},
	}
)

func letterCombinations(digits string) []string {
	ans := make([]string, 0)
	if len(digits) == 0 {
		return ans
	}
	letters := make([]byte, len(digits))
	ans = combine(digits, 0, letters, ans)
	return ans
}

func combine(digits string, index int, letters []byte, ans []string) []string {
	isEnd := (index == len(letters)-1)
	for _, ch := range numToLetters[digits[index]] {
		letters[index] = ch
		if isEnd {
			ans = append(ans, string(letters))
		} else {
			ans = combine(digits, index+1, letters, ans)
		}
	}
	return ans
}
