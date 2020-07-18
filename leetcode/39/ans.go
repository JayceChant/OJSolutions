package main

import "sort"

// 0ms, 100%
// 3.2MB, 71.43%
func combinationSum(candidates []int, target int) [][]int {
	answers := make([][]int, 0)
	size := len(candidates)
	if size == 0 {
		return answers
	}

	ansCand := make([]int, 0)
	sort.Ints(candidates)
	return combine(candidates, target, 0, size, 0, ansCand, answers)
}

func combine(candidates []int, target int, i int, size int, count int, ansCand []int, answers [][]int) [][]int {
	num := candidates[i]
	hasNext := (i+1 < size)
	for {
		if hasNext {
			hasNext = (candidates[i+1] <= target)
			if hasNext {
				answers = combine(candidates, target, i+1, size, count, ansCand, answers)
			}
		}

		target -= num
		if target < 0 {
			break
		}

		ansCand = append(ansCand, num)
		count++
		if target == 0 {
			ans := make([]int, count)
			copy(ans, ansCand)
			answers = append(answers, ans)
			break
		}
	}
	return answers
}

// try DP? How to efficiently remove duplicates?
