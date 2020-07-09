package main

import "sort"

// 能够重排成等差数列？
// 2 <= arr.length <= 1000
// -10^6 <= arr[i] <= 10^6
func canMakeArithmeticProgression(arr []int) bool {
	sort.Ints(arr)
	diff := arr[1] - arr[0]
	for i := 2; i< len(arr);i++{
		if arr[i] - arr[i-1] != diff {
			return false
		}
	}
	return true
}