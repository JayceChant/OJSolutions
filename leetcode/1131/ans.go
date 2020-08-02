package main

import "math"

// 56ms
// 6.5MB
func maxAbsValExpr(arr1 []int, arr2 []int) int {
	size := len(arr1)
	maxpp, maxpn, maxnp, maxnn := math.MinInt32, math.MinInt32, math.MinInt32, math.MinInt32
	minpp, minpn, minnp, minnn := math.MaxInt32, math.MaxInt32, math.MaxInt32, math.MaxInt32
	for i := 0; i < size; i++ {
		// ++
		num := arr1[i] + arr2[i] + i
		if maxpp < num {
			maxpp = num
		}
		if minpp > num {
			minpp = num
		}
		// +-
		num = arr1[i] + arr2[i] - i
		if maxpn < num {
			maxpn = num
		}
		if minpn > num {
			minpn = num
		}
		// -+
		num = arr1[i] - arr2[i] + i
		if maxnp < num {
			maxnp = num
		}
		if minnp > num {
			minnp = num
		}
		// --
		num = arr1[i] - arr2[i] - i
		if maxnn < num {
			maxnn = num
		}
		if minnn > num {
			minnn = num
		}
	}
	max := maxpp - minpp
	diff := maxpn - minpn
	if max < diff {
		max = diff
	}
	diff = maxnp - minnp
	if max < diff {
		max = diff
	}
	diff = maxnn - minnn
	if max < diff {
		max = diff
	}
	return max
}
