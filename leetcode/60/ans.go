package main

import "strings"

// 0ms
// 2MB
func getPermutation(n int, k int) string {
	used := make([]bool, 9)
	sb := strings.Builder{}
	k-- // start from 0
	for rest := n - 1; rest >= 0; rest-- {
		digit := byte(k / fact[rest])
		k %= fact[rest]
		for i := byte(0); i <= digit; i++ {
			if used[i] {
				digit++
			}
		}
		sb.WriteByte('1' + digit)
		used[digit] = true
	}
	return sb.String()
}

var (
	fact = [...]int{
		1,
		1,
		2,
		6,
		24,
		120,
		720,
		5040,
		40320,
	}
)
