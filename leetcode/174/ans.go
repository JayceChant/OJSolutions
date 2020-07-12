package main

// binary search
func calculateMinimumHPBS(dungeon [][]int) int {
	rowNum := len(dungeon)
	colNum := len(dungeon[0])

	maxHPLastRow := make([]int, colNum)
	maxHPThisRow := make([]int, colNum)

	// 1st row special →
	maxHPThisRow[0] = dungeon[0][0]
	lowestHP := maxHPThisRow[0]
	for c := 1; c < colNum; c++ {
		maxHPThisRow[c] = maxHPThisRow[c-1] + dungeon[0][c]
		lowestHP = min(lowestHP, maxHPThisRow[c])
	}

	for r := 1; r < rowNum; r++ {
		// swap dp row to reuse
		maxHPLastRow, maxHPThisRow = maxHPThisRow, maxHPLastRow
		// 1st col special ↓
		maxHPThisRow[0] = maxHPLastRow[0] + dungeon[r][0]
		lowestHP = min(lowestHP, maxHPThisRow[0])

		for c := 1; c < colNum; c++ {
			// from the direction with higher HP
			maxHPThisRow[c] = max(maxHPThisRow[c-1], maxHPLastRow[c]) + dungeon[r][c]
			lowestHP = min(lowestHP, maxHPThisRow[c])
		}
	}

	// (low, high]
	low := 0
	high := max(1, 1-lowestHP)

hpLoop:
	for low+1 < high {
		tryHP := (low + high) / 2

		// 1st row special →
		maxHPThisRow[0] = tryHP + dungeon[0][0]
		if maxHPThisRow[0] < 1 {
			low = tryHP
			continue hpLoop
		}

		for c := 1; c < colNum; c++ {
			if maxHPThisRow[c-1] > 0 {
				maxHPThisRow[c] = maxHPThisRow[c-1] + dungeon[0][c]
			} else {
				maxHPThisRow[c] = 0 // unreachable
			}
		}

		for r := 1; r < rowNum; r++ {
			// swap dp row to reuse
			maxHPLastRow, maxHPThisRow = maxHPThisRow, maxHPLastRow
			// survived flag to early prune
			survived := false
			// 1st col special ↓
			if maxHPLastRow[0] > 0 {
				maxHPThisRow[0] = maxHPLastRow[0] + dungeon[r][0]
				if maxHPThisRow[0] > 0 {
					survived = true
				}
			} else {
				maxHPThisRow[0] = 0
			}

			for c := 1; c < colNum; c++ {
				// from the direction with higher HP
				preHP := max(maxHPThisRow[c-1], maxHPLastRow[c])
				if preHP > 0 {
					maxHPThisRow[c] = preHP + dungeon[r][c]
					if maxHPThisRow[c] > 0 {
						survived = true
					}
				} else {
					maxHPThisRow[c] = 0
				}
			}
			if !survived {
				low = tryHP
				continue hpLoop
			}
		}

		if maxHPThisRow[colNum-1] > 0 {
			high = tryHP
		} else {
			low = tryHP
		}
	}
	return high
}

func calculateMinimumHPBS2(dungeon [][]int) int {
	rowNum := len(dungeon)
	colNum := len(dungeon[0])

	maxHPLastRow := make([]int, colNum)
	maxHPThisRow := make([]int, colNum)

	// (low, high]
	low := 0
	high := 0
	tryHP := 1
hpLoop:
	for {
		if low+1 < high {
			tryHP = (low + high) / 2
		} else if high == 0 {
			tryHP *= 2
		} else {
			break
		}

		// 1st row special →
		maxHPThisRow[0] = tryHP + dungeon[0][0]
		if maxHPThisRow[0] < 1 {
			low = tryHP
			continue hpLoop
		}

		for c := 1; c < colNum; c++ {
			if maxHPThisRow[c-1] > 0 {
				maxHPThisRow[c] = maxHPThisRow[c-1] + dungeon[0][c]
			} else {
				maxHPThisRow[c] = 0 // unreachable
			}
		}

		for r := 1; r < rowNum; r++ {
			// swap dp row to reuse
			maxHPLastRow, maxHPThisRow = maxHPThisRow, maxHPLastRow
			// survived flag to early prune
			survived := false
			// 1st col special ↓
			if maxHPLastRow[0] > 0 {
				maxHPThisRow[0] = maxHPLastRow[0] + dungeon[r][0]
				if maxHPThisRow[0] > 0 {
					survived = true
				}
			} else {
				maxHPThisRow[0] = 0
			}

			for c := 1; c < colNum; c++ {
				// from the direction with higher HP
				preHP := max(maxHPThisRow[c-1], maxHPLastRow[c])
				if preHP > 0 {
					maxHPThisRow[c] = preHP + dungeon[r][c]
					if maxHPThisRow[c] > 0 {
						survived = true
					}
				} else {
					maxHPThisRow[c] = 0
				}
			}
			if !survived {
				low = tryHP
				continue hpLoop
			}
		}

		if maxHPThisRow[colNum-1] > 0 {
			high = tryHP
		} else {
			low = tryHP
		}
	}
	return high
}

// (0,0) -> ed is not equivalent to ed -> (0,0)
// 8 ms, faster than 19.28% (min 4ms)
// 3.7 MB, less than 88.89%
func calculateMinimumHPDP(dungeon [][]int) int {
	lastRow := len(dungeon) - 1
	lastCol := len(dungeon[0]) - 1

	// 1st row
	dungeon[lastRow][lastCol] = 1 - dungeon[lastRow][lastCol]
	if dungeon[lastRow][lastCol] < 1 {
		dungeon[lastRow][lastCol] = 1
	}
	for c := lastCol - 1; c >= 0; c-- {
		dungeon[lastRow][c] = dungeon[lastRow][c+1] - dungeon[lastRow][c]
		if dungeon[lastRow][c] < 1 {
			dungeon[lastRow][c] = 1
		}
	}
	for r := lastRow - 1; r >= 0; r-- {
		dungeon[r][lastCol] = dungeon[r+1][lastCol] - dungeon[r][lastCol]
		if dungeon[r][lastCol] < 1 {
			dungeon[r][lastCol] = 1
		}

		for c := lastCol - 1; c >= 0; c-- {
			dungeon[r][c] = min(dungeon[r][c+1], dungeon[r+1][c]) - dungeon[r][c]
			if dungeon[r][c] < 1 {
				dungeon[r][c] = 1
			}
		}
	}
	//fmt.Println(dungeon)
	if dungeon[0][0] < 1 {
		return 1
	}
	return dungeon[0][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
