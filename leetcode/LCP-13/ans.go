package main

import (
	"container/list"
	"math"
)

// 528ms (min 84ms)
// 40.9MB
func minimalSteps(maze []string) int {
	rowNum = len(maze)
	colNum = len(maze[0])
	mazeB = maze

	var rSt, cSt, rEd, cEd, countM int
	for r := 0; r < rowNum; r++ {
		for c := 0; c < colNum; c++ {
			switch mazeB[r][c] {
			case 'S':
				rSt = r
				cSt = c
			case 'T':
				rEd = r
				cEd = c
			case 'M':
				countM++
			}
		}
	}

	if countM == 0 {
		path := bfsPath(rSt, cSt, 'T')
		if len(path) == 0 {
			return -1
		}
		return path[0][2]
	}

	// all M accessible from ed
	msFromEd := bfsPath(rEd, cEd, 'M')
	// fmt.Println(msFromEd)
	if len(msFromEd) < countM {
		// some M(s) is/are not accessible
		return -1
	}

	// step from st to all O
	osFromSt := bfsPath(rSt, cSt, 'O')
	if len(osFromSt) == 0 {
		// no O is accessible
		return -1
	}

	// step from each M to each O
	stepMToOs := make([](map[int]int), countM)
	for i, m := range msFromEd {
		stepMToOs[i] = map[int]int{}
		mToOs := bfsPath(m[0], m[1], 'O')
		for _, o := range mToOs {
			oid := o[0]*100 + o[1]
			stepMToOs[i][oid] = o[2]
		}
	}
	// fmt.Println(stepMToOs)

	// min step from st to each M by O
	stepStToMsByO := make([]int, countM)
	// min step from each M to other M by O
	stepMToMsByO := make([][]int, countM)
	for i := 0; i < countM; i++ {
		stepMToMsByO[i] = make([]int, countM)
	}

	for i, oToStep := range stepMToOs {
		stepStToMsByO[i] = math.MaxInt32
		for _, o := range osFromSt {
			oid := o[0]*100 + o[1]
			step, ok := oToStep[oid]
			if !ok {
				// an O from st can't access an M from ed
				return -1
			}
			step += o[2]
			if stepStToMsByO[i] > step {
				stepStToMsByO[i] = step
			}
		}

		for i2 := i + 1; i2 < countM; i2++ {
			stepMToMsByO[i][i2] = math.MaxInt32
			stepMToMsByO[i2][i] = math.MaxInt32
			for oid, step := range oToStep {
				step += stepMToOs[i2][oid]
				if stepMToMsByO[i][i2] > step {
					stepMToMsByO[i][i2] = step
					stepMToMsByO[i2][i] = step
				}
			}
		}
	}
	// fmt.Println(stepStToMsByO)
	// fmt.Println(stepMToMsByO)

	// dp[i<<countM + mask] means Ms in mask are triggered
	// mask is a bitmap that bit i is 1 means M[i] is triggered
	dp := map[int]int{}
	for i, step := range stepStToMsByO {
		dp[i<<countM+1<<i] = step
	}

	fullMask := 1<<countM - 1
	for mask := 1; mask < fullMask; mask++ {
		for curr := 0; 1<<curr <= mask; curr++ {
			bitC := 1 << curr
			//fmt.Println("bitCurr", mask, curr, bitC)
			if mask&bitC == 0 {
				continue // curr not in mask
			}

			for next := 0; next < countM; next++ {
				bitN := 1 << next
				//fmt.Println("bitNext", mask, next, bitN)
				if mask&bitN != 0 {
					continue // next already in mask
				}

				step := dp[curr<<countM+mask] + stepMToMsByO[curr][next]
				key := next<<countM + (mask | bitN)
				//fmt.Println(key)
				if minStep, ok := dp[key]; !ok || minStep > step {
					dp[key] = step
				}
			}
		}
	}
	//fmt.Println(dp)

	minStep := math.MaxInt32
	for last := 0; last < countM; last++ {
		step := dp[last<<countM+fullMask] + msFromEd[last][2]
		if minStep > step {
			minStep = step
		}
	}
	return minStep
}

var (
	rowNum, colNum int
	mazeB          []string
)

func bfsPath(row, col int, target byte) (points [][3]int) {
	visited := make([][]bool, rowNum)
	for r := 0; r < rowNum; r++ {
		visited[r] = make([]bool, colNum)
	}

	queue := list.New()
	queue.PushBack([3]int{row, col, 0})
	visited[row][col] = true
	var step int
	for queue.Len() > 0 {
		e := queue.Front()
		queue.Remove(e)
		point := e.Value.([3]int)
		row, col, step = point[0], point[1], point[2]+1

		// up
		if row > 0 && !visited[row-1][col] {
			if mazeB[row-1][col] == target {
				points = append(points, [3]int{row - 1, col, step})
			}
			if mazeB[row-1][col] != '#' {
				queue.PushBack([3]int{row - 1, col, step})
				visited[row-1][col] = true
			}
		}
		// down
		if row < rowNum-1 && !visited[row+1][col] {
			if mazeB[row+1][col] == target {
				points = append(points, [3]int{row + 1, col, step})
			}
			if mazeB[row+1][col] != '#' {
				queue.PushBack([3]int{row + 1, col, step})
				visited[row+1][col] = true
			}
		}
		// left
		if col > 0 && !visited[row][col-1] {
			if mazeB[row][col-1] == target {
				points = append(points, [3]int{row, col - 1, step})
			}
			if mazeB[row][col-1] != '#' {
				queue.PushBack([3]int{row, col - 1, step})
				visited[row][col-1] = true
			}
		}
		// right
		if col < colNum-1 && !visited[row][col+1] {
			if mazeB[row][col+1] == target {
				points = append(points, [3]int{row, col + 1, step})
			}
			if mazeB[row][col+1] != '#' {
				queue.PushBack([3]int{row, col + 1, step})
				visited[row][col+1] = true
			}
		}
	}
	// no path
	return
}

// a wrong greedy solution
// taking the nearest route at every step
// can lead to a possible long detour globally
// fail case:
// var (
// 	mazeFail = []string{
// 		"S.O",
// 		"...",
// 		"...",
// 		"O..",
// 		"M..",
// 		"T..",
// 	}
// )

// func minimalSteps(maze []string) int {
// 	rowNum = len(maze)
// 	colNum = len(maze[0])
// 	mazeB = make([][]byte, rowNum)

// 	var rSt, cSt, countM int
// 	for r := 0; r < rowNum; r++ {
// 		mazeB[r] = []byte(maze[r])
// 		for c := 0; c < colNum; c++ {
// 			switch mazeB[r][c] {
// 			case 'S':
// 				rSt = r
// 				cSt = c
// 			case 'M':
// 				countM++
// 			}
// 		}
// 	}

// 	var stepSum, step int
// 	for countM > 0 {
// 		rSt, cSt, step = bfsPath(rSt, cSt, 'O')
// 		if step == -1 {
// 			return -1
// 		}
// 		stepSum += step

// 		rSt, cSt, step = bfsPath(rSt, cSt, 'M')
// 		if step == -1 {
// 			return -1
// 		}
// 		stepSum += step
// 		mazeB[rSt][cSt] = '.'
// 		countM--
// 	}
// 	_, _, step = bfsPath(rSt, cSt, 'T')
// 	if step == -1 {
// 		return -1
// 	}
// 	stepSum += step

// 	return stepSum
// }
