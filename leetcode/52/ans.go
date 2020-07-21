package main

var (
	cols []bool
	// diagonal left-top to right-bottom
	// order from left-bottom to right-top: index = row + col
	diag []bool
	// diagonal left-bottom to right-top
	// order from left-top to right-bottom: index = n - row + col - 1
	diag2 []bool
)

// constraint programming + back-tracking
// 0ms, 100%
// 2.1MB, 33.33%
func totalNQueens(n int) int {
	if n == 0 {
		return 0
	}

	cols = make([]bool, n)
	diag = make([]bool, 2*n-1)
	diag2 = make([]bool, 2*n-1)

	return putQueen(n, 0)
}

func putQueen(n int, row int) int {
	count := 0
	for col := 0; col < n; col++ {
		if cols[col] || diag[row+col] || diag2[n-row+col-1] {
			continue
		}

		if row == n-1 {
			count++
		} else {
			cols[col] = true
			diag[row+col] = true
			diag2[n-row+col-1] = true
			count += putQueen(n, row+1)
			diag2[n-row+col-1] = false
			diag[row+col] = false
			cols[col] = false
		}
	}
	return count
}
