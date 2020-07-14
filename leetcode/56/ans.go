package main

import (
	"math"
	"sort"
)

// 12ms, 82.80% (min 4ms)
// 4.6MB, 100%
func merge(intervals [][]int) [][]int {
	ans := make([][]int, 0)
	sort.Sort(intervalSlice(intervals))

	last := []int{math.MinInt32, math.MinInt32}
	for i := 0; i < len(intervals); i++ {
		if last[1] < intervals[i][0] { // no intersection
			last = intervals[i]
			ans = append(ans, last)
			continue
		}

		if last[1] < intervals[i][1] { // intersected, merge
			last[1] = intervals[i][1]
		}
		// else, the latter is included, do nothing
	}
	return ans
}

var _ sort.Interface = intervalSlice(nil)

type intervalSlice [][]int

func (is intervalSlice) Len() int { return len(is) }

func (is intervalSlice) Less(i, j int) bool { return is[i][0] < is[j][0] }

func (is intervalSlice) Swap(i, j int) { is[i], is[j] = is[j], is[i] }
