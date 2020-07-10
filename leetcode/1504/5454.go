package main

// 统计全 1 子矩形
// 1 <= rows <= 150
// 1 <= columns <= 150
// 0 <= mat[i][j] <= 1
func numSubmat(mat [][]int) int {
	count := 0
	row := len(mat)
	col := len(mat[0])

	mymat := make([][]int, row)
	for i := 0; i < row; i++ {
		mymat[i] = make([]int, col)
		copy(mymat[i], mat[i])
	}
	// fmt.Println(mymat)

	for height := 1; height <= row; height++ {
		for r := 0; r <= row-height; r++ {
			length := 0
			for c := 0; c < col; c++ {
				if mymat[r][c] == 1 {
					length++
					count += length
				} else {
					length = 0
				}
			}
		}

		// compress
		for r := 0; r < row-height; r++ {
			for c := 0; c < col; c++ {
				if mymat[r][c] == 1 && mat[r+height][c] == 1 {
					mymat[r][c] = 1
				} else {
					mymat[r][c] = 0
				}
			}
		}
	}
	return count
}
