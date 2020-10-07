package main

import "sort"

// 0ms
// 2MB
func splitArraySameAverage(A []int) bool {
	size = len(A)
	sum := 0
	for _, a := range A {
		sum += a
	}
	d := gcd(sum, size)
	if d == 1 {
		return false
	}
	base := size / d
	baseSum := sum / d
	sort.Ints(A)
	ga = A
	halfd := d / 2

	for i := 1; i <= halfd; i++ {
		if nSum(base*i, baseSum*i, 0, 0) {
			return true
		}
	}
	return false
}

var (
	size int
	ga   []int
)

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func nSum(n int, target int, i int, st int) bool {
	if n == 1 {
		x := sort.SearchInts(ga, target)
		return ga[x] == target
	}

	if i == n-2 { // two sum
		l := st
		r := size - 1
		for l < r {
			sum := ga[l] + ga[r]
			if sum == target {
				return true
			}
			if sum < target {
				l++
			} else {
				r--
			}
		}
		return false
	}

	for j := st; size-j >= n-i; j++ {
		if j > st && ga[j] == ga[j-1] {
			continue
		}
		if nSum(n, target-ga[j], i+1, j+1) {
			return true
		}
	}
	return false
}
