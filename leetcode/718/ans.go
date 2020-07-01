package main

import (
	"fmt"
	"math"
)

// 最长公共子数组（注意不是子序列）

// O(NM), 40 ms, faster than 84.31% (min 20ms)
// O(NM), 3.4 MB, less than 92.86% (min 3376kB)
func findLengthDP(A []int, B []int) int {
	sizeA := len(A)
	sizeB := len(B)
	if sizeA < sizeB {
		A, B, sizeA, sizeB = B, A, sizeB, sizeA
	}
	// size <= 1000
	max := uint16(0)
	// dp index starts from 1
	dp := make([][]uint16, 2)
	for i := 0; i < 2; i++ {
		dp[i] = make([]uint16, sizeB+1)
	}
	for i := 0; i < sizeA; i++ {
		for j := 0; j < sizeB; j++ {
			if A[i] == B[j] {
				dp[1][j+1] = dp[0][j] + 1
				if dp[1][j+1] > max {
					max = dp[1][j+1]
				}
			} else {
				dp[1][j+1] = 0
			}
		}
		dp[0], dp[1] = dp[1], dp[0]
	}
	return int(max)
}

// O((N+M)min(N, M)), 32 ms, faster than 86.27%
// O(1), 3.4 MB, less than 100.00%
func findLengthOffset(A []int, B []int) int {
	max := 0
	for t := 0; t < 2; t++ {
		for offset := 0; offset < len(A); offset++ {
			length := 0
			end := len(A) - offset
			if len(B) < end {
				end = len(B)
			}
			for i := 0; i < end; i++ {
				if A[i+offset] == B[i] {
					length++
					if length > max {
						max = length
					}
				} else {
					length = 0
				}
			}
		}
		A, B = B, A
	}
	return max
}

const (
	base    = 101
	divisor = 1000000009
)

func findLengthHash(A []int, B []int) int {
	sizeA := len(A)
	sizeB := len(B)
	// max in [low, high)
	low := 1
	high := sizeA
	if sizeB < high {
		high = sizeB
	}
	high++
LengthLoop:
	for low < high {
		length := (low + high) / 2
		mult := quickPow(base, length-1)

		existed := make(map[int]bool, 0)
		hash := 0
		for i := 0; i < length; i++ {
			hash = (hash*base + A[i]) % divisor
		}
		existed[hash] = true
		for i := length; i < sizeA; i++ {
			hash = ((hash-A[i-length]*mult%divisor+divisor)%divisor*base + A[i]) % divisor
			existed[hash] = true
		}
		hash = 0
		for i := 0; i < length; i++ {
			hash = (hash*base + B[i]) % divisor
		}
		if existed[hash] {
			low = length + 1
			continue
		}
		for i := length; i < sizeB; i++ {
			hash = ((hash-B[i-length]*mult%divisor+divisor)%divisor*base + B[i]) % divisor
			if existed[hash] {
				low = length + 1
				continue LengthLoop
			}
		}
		high = length
	}
	return low - 1
}

func quickPow(x int, n int) int {
	ans := 1
	for n > 0 {
		if n&1 != 0 {
			ans *= x
		}
		n >>= 1
		x *= x
	}
	return ans
}

// max PN <= x
// 32749 < MaxInt16
// 65521 < MaxUint16
// MaxInt32
func primeNum(end uint64) []uint64 {
	ans := []uint64{2, 3, 5, 7, 11}
	for num := uint64(13); num < end; num += 2 {
		sqrt := uint64(math.Sqrt(float64(num))) + 1
		for i := 0; i < len(ans); i++ {
			if ans[i] >= sqrt {
				ans = append(ans, num)
				fmt.Println(num)
				break
			}
			if num%ans[i] == 0 {
				break
			}
		}
	}
	return ans
}

func main() {
	primeNum(math.MaxInt32)
}
