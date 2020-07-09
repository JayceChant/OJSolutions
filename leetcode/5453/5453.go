package main

// 最后一只蚂蚁掉落的时刻
// 1 <= n <= 10^4
// 0 <= left.length <= n + 1
// 0 <= left[i] <= n
// 0 <= right.length <= n + 1
// 0 <= right[i] <= n
// 1 <= left.length + right.length <= n + 1
// left 和 right 中的所有值都是唯一的，并且每个值 只能出现在二者之一 中。

func getLastMoment(n int, left []int, right []int) int {
	max := 0
	for _, x := range left {
		if x > max {
			max = x
		}
	}

	for _, x := range right {
		diff := n - x
		if diff > max {
			max = diff
		}
	}
	return max
}
