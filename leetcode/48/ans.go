package main

// 0ms
// 2.2MB
func rotate(matrix [][]int) {
	n := len(matrix)
	// transposition
	for r := 0; r < n-1; r++ {
		for c := r + 1; c < n; c++ {
			matrix[r][c], matrix[c][r] = matrix[c][r], matrix[r][c]
		}
	}
	// lef-right mirror
	for r := 0; r < n; r++ {
		for c := 0; c < n/2; c++ {
			matrix[r][c], matrix[r][n-c-1] = matrix[r][n-c-1], matrix[r][c]
		}
	}
}
