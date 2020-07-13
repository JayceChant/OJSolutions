package main

var (
	strToInt = map[string]int{
		"M":  1000,
		"CM": 900,
		"D":  500,
		"CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
)

// 12ms 49.14% (min 0ms)
// 3.1MB 100%
func romanToInt1(s string) int {
	ans := 0
	i := 0
	for i < len(s)-1 {
		if d := strToInt[s[i:i+2]]; d > 0 {
			ans += d
			i += 2
		} else {
			ans += strToInt[s[i:i+1]]
			i++
		}
	}
	if i < len(s) {
		ans += strToInt[s[i:i+1]]
	}
	return ans
}

var (
	charToInt = map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
)

// 4ms, 94.60% (min 0ms)
// 3.1MB, 40% (mine 3076, min 3072)
func romanToInt(s string) int {
	right := 0
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		cur := charToInt[s[i]]
		if cur >= right {
			ans += cur
		} else {
			ans -= cur
		}
		right = cur
	}
	return ans
}
