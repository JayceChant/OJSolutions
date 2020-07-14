package main

func minimumTotal(triangle [][]int) int {
	r := len(triangle) - 2
	for r >= 0 {
		for i := 0; i <= r; i++ {
			if triangle[r+1][i] <= triangle[r+1][i+1] {
				triangle[r][i] += triangle[r+1][i]
			} else {
				triangle[r][i] += triangle[r+1][i+1]
			}
		}
		r--
	}
	return triangle[0][0]
}
