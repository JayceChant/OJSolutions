package main

// 4ms, 39.82
// 2.1MB, 29.41%
func divisorGameMemo(N int) bool {
	if N == 1 {
		return false
	}

	if res, exist := cache[N]; exist {
		return res
	}

	for x := N - 1; x > 0; x-- {
		if N%x == 0 && !divisorGame(N-x) {
			cache[N] = true
			return true
		}
	}
	cache[N] = false
	return false
}

var (
	cache = map[int]bool{}
)

// 0ms, 100%
// 1.9MB, 94.12%
func divisorGame(N int) bool {
	return N&1 == 0
}
