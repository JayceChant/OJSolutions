package main

const (
	modBase = 1000000007
)

func numSub(s string) int {
	count := 0

	length := 0
	for _, ch := range s {
		if ch == '1' {
			length++
			count = (count + length) % modBase
		} else {
			length = 0
		}
	}

	return count
}
