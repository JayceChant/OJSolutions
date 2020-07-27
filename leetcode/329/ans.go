package main

import "sort"

// 60ms (min 24ms)
// 6.3MB (min 6288)
func longestIncreasingPath(matrix [][]int) int {
	rowNum = len(matrix)
	if rowNum == 0 {
		return 0
	}
	colNum = len(matrix[0])
	if colNum == 0 {
		return 0
	}
	mat = matrix
	memo = make([][]int, rowNum)
	maxLen = 0
	for r := 0; r < rowNum; r++ {
		memo[r] = make([]int, len(matrix[0]))
	}

	for r := 0; r < rowNum; r++ {
		for c := 0; c < colNum; c++ {
			if (r > 0 && matrix[r][c] > matrix[r-1][c]) ||
				(c > 0 && matrix[r][c] > matrix[r][c-1]) ||
				(r < rowNum-1 && matrix[r][c] > matrix[r+1][c]) ||
				(c < colNum-1 && matrix[r][c] > matrix[r][c+1]) {
				continue
			}
			climb(r, c, 1)
		}
	}

	return maxLen
}

var (
	mat    [][]int
	rowNum int
	colNum int
	memo   [][]int
	maxLen int
)

func climb(row, col, length int) {
	memo[row][col] = length
	length++
	moved := false
	if row > 0 && mat[row-1][col] > mat[row][col] && memo[row-1][col] < length {
		climb(row-1, col, length)
		moved = true
	}

	if col > 0 && mat[row][col-1] > mat[row][col] && memo[row][col-1] < length {
		climb(row, col-1, length)
		moved = true
	}

	if row < rowNum-1 && mat[row+1][col] > mat[row][col] && memo[row+1][col] < length {
		climb(row+1, col, length)
		moved = true
	}

	if col < colNum-1 && mat[row][col+1] > mat[row][col] && memo[row][col+1] < length {
		climb(row, col+1, length)
		moved = true
	}

	length--
	if !moved && maxLen < length {
		maxLen = length
	}
}

// 52ms
// 6.8MB
func longestIncreasingPathSort(matrix [][]int) int {
	rowNum := len(matrix)
	if rowNum == 0 {
		return 0
	}
	colNum := len(matrix[0])
	if colNum == 0 {
		return 0
	}

	max := 0

	points := make([][]int, 0)
	for r, row := range matrix {
		for c, num := range row {
			points = append(points, []int{num, r + 1, c + 1})
		}
	}

	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0]
	})

	memo := make([][]int, rowNum+2)
	for r := 0; r < rowNum+2; r++ {
		memo[r] = make([]int, colNum+2)
	}

	for _, p := range points {
		length := 0
		num, r, c := p[0], p[1], p[2]
		if length < memo[r-1][c] && num != matrix[r-2][c-1] {
			length = memo[r-1][c]
		}
		if length < memo[r+1][c] && num != matrix[r][c-1] {
			length = memo[r+1][c]
		}
		if length < memo[r][c-1] && num != matrix[r-1][c-2] {
			length = memo[r][c-1]
		}
		if length < memo[r][c+1] && num != matrix[r-1][c] {
			length = memo[r][c+1]
		}

		length++
		memo[r][c] = length

		if max < length {
			max = length
		}
	}

	return max
}
