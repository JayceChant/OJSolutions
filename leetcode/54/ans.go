package main

// 0ms
// 2MB
func spiralOrder(matrix [][]int) (ans []int) {
	m := len(matrix)
	if m == 0 {
		return
	}
	n := len(matrix[0])

	t := 0
	row := 0
	col := -1

	for n > 0 {
		for i := 0; i < n; i++ {
			row += directions[t%4][0]
			col += directions[t%4][1]
			ans = append(ans, matrix[row][col])
		}
		t++
		n, m = m-1, n
	}
	return
}

var (
	directions = [4][2]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}
)
