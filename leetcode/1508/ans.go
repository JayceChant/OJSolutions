package main

import "sort"

const (
	modBase = 1000000007
)

func rangeSum(nums []int, n int, left int, right int) int {
	sums := make([]int, 0, n*(n+1)/2)
	for st := 0; st < n; st++ {
		sum := 0
		for i := st; i < n; i++ {
			sum += nums[i]
			sums = append(sums, sum)
		}
	}
	sort.Ints(sums)
	sum := 0
	for i := left - 1; i < right; i++ {
		sum = (sum + sums[i]) % modBase
	}
	return sum
}
