package main

// 0ms
// 2.1MB
func numJewelsInStones(J string, S string) int {
	hash := map[byte]bool{}
	for i := range J {
		hash[J[i]] = true
	}

	count := 0
	for i := range S {
		if hash[S[i]] {
			count++
		}
	}
	return count
}
