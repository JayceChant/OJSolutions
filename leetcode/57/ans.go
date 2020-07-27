package main

// 12ms, 28.7%
// 4.7MB, 100%
func insert(intervals [][]int, newInterval []int) [][]int {
	size := len(intervals)
	ans := make([][]int, 0)

	i := 0
	for i < size && intervals[i][1] < newInterval[0] {
		ans = append(ans, intervals[i])
		i++
	}

	if i < size && intervals[i][0] < newInterval[0] {
		// merge left of first intersection
		newInterval[0] = intervals[i][0]
	}

	for i < size && intervals[i][1] <= newInterval[1] {
		i++
	}

	if i < size && intervals[i][0] <= newInterval[1] {
		// merge right of last intersection
		newInterval[1] = intervals[i][1]
		i++
	}

	ans = append(ans, newInterval)

	for i < size {
		ans = append(ans, intervals[i])
		i++
	}

	return ans
}

// 12ms, 28.7%
// 4.7MB, 100%
func insert2(intervals [][]int, newInterval []int) [][]int {
	ans := make([][]int, 0)

	inserted := false
	for _, interval := range intervals {
		if interval[1] < newInterval[0] {
			ans = append(ans, interval)
			continue
		}

		if !inserted {
			if interval[0] < newInterval[0] {
				newInterval[0] = interval[0]
			}
			ans = append(ans, newInterval)
			inserted = true
		}

		if newInterval[1] < interval[0] {
			ans = append(ans, interval)
			continue
		}

		if newInterval[1] < interval[1] {
			newInterval[1] = interval[1]
		}
	}

	if !inserted {
		ans = append(ans, newInterval)
	}

	return ans
}
