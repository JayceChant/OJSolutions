package main

import "strings"

var (
	queens []int
	rowStr string
	cols   []bool
	// diagonal left-top to right-bottom
	// order from left-bottom to right-top: index = row + col
	diag []bool
	// diagonal left-bottom to right-top
	// order from left-top to right-bottom: index = n - row + col - 1
	diag2 []bool
)

// constraint programming + back-tracking
// 4ms, 91.82%
// 3.2MB, 96%
func solveNQueens(n int) [][]string {
	if n == 0 {
		return [][]string{{}}
	}

	queens = make([]int, n)
	rowStr = "Q" + strings.Repeat(".", n-1)

	cols = make([]bool, n)
	diag = make([]bool, 2*n-1)
	diag2 = make([]bool, 2*n-1)

	solutions := make([][]string, 0)
	return putQueen(n, 0, solutions)
}

func putQueen(n int, row int, solutions [][]string) [][]string {
	for col := 0; col < n; col++ {
		if cols[col] || diag[row+col] || diag2[n-row+col-1] {
			continue
		}

		queens[row] = col

		if row == n-1 {
			solution := make([]string, n)
			for r := 0; r < n; r++ {
				c := queens[r]
				solution[r] = rowStr[1:c+1] + rowStr[:n-c]
			}
			solutions = append(solutions, solution)
		} else {
			cols[col] = true
			diag[row+col] = true
			diag2[n-row+col-1] = true
			solutions = putQueen(n, row+1, solutions)
			diag2[n-row+col-1] = false
			diag[row+col] = false
			cols[col] = false
		}
	}
	return solutions
}
