package main

import (
	"math"
	"sort"
)

// 1 <= arr.length <= 10^5
// 1 <= arr[i] <= 10^6
// 0 <= target <= 10^7

// O(N) (3*20N)
// 136ms, 80%
// 9.1MB, 100%
func closestToTarget(arr []int, target int) int {
	min := math.MaxInt32
	size := len(arr)

	andProducts := make([]int, 0)

	for r := 0; r < size; r++ {
		for i := 0; i < len(andProducts); i++ {
			andProducts[i] &= arr[r]
		}
		andProducts = append(andProducts, arr[r])
		sort.Ints(andProducts)
		andProducts = dedup(andProducts)

		for _, ap := range andProducts {
			diff := myAbs(ap - target)
			if diff == 0 {
				return 0
			}

			if min > diff {
				min = diff
			} else if ap > target {
				break
			}
		}
	}

	return min
}

func myAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func dedup(sorted []int) []int {
	newSize := 1

	for i := 1; i < len(sorted); i++ {
		if sorted[i] == sorted[i-1] {
			continue
		}

		if newSize < i {
			sorted[newSize] = sorted[i]
		}
		newSize++
	}

	return sorted[:newSize]
}

// sample f()
// func f(arr []int, l, r int) int {
// 	if r < l {
// 		return failVal
// 	}
// 	ans := arr[l]
// 	for i := l + 1; l <= r; i++ {
// 		ans &= arr[i]
// 	}
// 	return ans
// }
