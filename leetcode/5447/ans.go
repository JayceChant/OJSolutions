package main

import "math"

func winnerSquareGameBF(n int) bool {
	if n == 0 {
		return false
	}
	root := int(math.Sqrt(float64(n)))
	for i := 1; i <= root; i++ {
		if !winnerSquareGameBF(n - i*i) {
			return true
		}
	}
	return false
}

var cache = map[int]bool{}

func winnerSquareGame(n int) bool {
	if n == 0 {
		return false
	}
	if win, ok := cache[n]; ok {
		return win
	}
	root := int(math.Sqrt(float64(n)))
	for i := root; i >= 1; i-- {
		if !winnerSquareGame(n - i*i) {
			cache[n] = true
			return true
		}
	}
	cache[n] = false
	return false
}
