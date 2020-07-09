package main

// 不同路径
// 0 ms, faster than 100.00%
// 2.7 MB, less than 18.82%
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	lastDp := make([]int, len(obstacleGrid[0]))
	curDp := make([]int, len(obstacleGrid[0]))
	lastDp[len(obstacleGrid[0])-1] = 1
	for row := len(obstacleGrid) - 1; row >= 0; row-- {
		for col := len(obstacleGrid[0]) - 1; col >= 0; col-- {
			curDp[col] = 0
			if obstacleGrid[row][col] == 0 {
				curDp[col] += lastDp[col]
				if col < len(obstacleGrid[0])-1 {
					curDp[col] += curDp[col+1]
				}
			}
		}
		curDp, lastDp = lastDp, curDp
	}
	//fmt.Println(curDp)
	//fmt.Println(lastDp)
	return lastDp[0]
}
