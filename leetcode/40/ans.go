package main

import "sort"

// 4ms, 63.14%
// 2.8MB, 50%
func combinationSum2(candidates []int, target int) [][]int {
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
	if i >= size || candidates[i] > target {
		return answers
	}

	// next indicates the index of next unique num
	next := i + 1
	for next < size && candidates[next] == candidates[next-1] {
		next++
	}

	answers = combine(candidates, target, next, size, count, ansCand, answers)

	num := candidates[i]
	for i < next {
		ansCand = append(ansCand, num)
		target -= num
		count++
		i++

		if target == 0 {
			ans := make([]int, count)
			copy(ans, ansCand)
			answers = append(answers, ans)
			break
		}

		answers = combine(candidates, target, next, size, count, ansCand, answers)
	}
	return answers
}
