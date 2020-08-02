package main

func minSwaps(grid [][]int) int {
	n := len(grid)
	reqZero := make([]int, n)
	for r := 0; r < n; r++ {
		c := n - 1
		for c >= 0 && grid[r][c] == 0 {
			c--
		}
		// (n - 1 - r) - (n - 1 - c), need - has
		reqZero[r] = c - r
	}

	count := 0
	for r := 0; r < n; r++ {
		if reqZero[r] <= 0 {
			continue
		}
		lastReq := reqZero[r]
		r2 := r + 1
		for r2 < n && reqZero[r2]+(r2-r) > 0 {
			lastReq, reqZero[r2] = reqZero[r2], lastReq-1
			r2++
		}
		if r2 == n {
			return -1
		}
		//reqZero[r] = reqZero[r2] + (r2 - r)
		reqZero[r2] = lastReq - 1
		count += (r2 - r)
	}
	return count
}
