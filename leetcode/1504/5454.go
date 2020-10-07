package main

// 统计全 1 子矩形
// 1 <= rows <= 150
// 1 <= columns <= 150
// 0 <= mat[i][j] <= 1
func numSubmat(mat [][]int) int {
	count := 0
	rowNum := len(mat)
	colNum := len(mat[0])

	mymat := make([]int, colNum)
	// fmt.Println(mymat)

	for rowSt := 0; rowSt < rowNum; rowSt++ {
		for rowEd := rowSt; rowEd < rowNum; rowEd++ {
			if rowEd == rowSt {
				copy(mymat, mat[rowSt])
			} else {
				for col := 0; col < colNum; col++ {
					if mat[rowEd][col] == 0 {
						mymat[col] = 0
					}
				}
			}

			length := 0
			for col := 0; col < colNum; col++ {
				if mymat[col] == 1 {
					length++
					count += length
				} else {
					length = 0
				}
			}
		}
	}
	return count
}

// 84 ms, faster than 23.53% (min 16ms)
// 6 MB, less than 100.00%
func numSubmat2(mat [][]int) int {
	count := 0
	rowNum := len(mat)
	colNum := len(mat[0])

	mymat := make([]bool, colNum)

	for rowSt := 0; rowSt < rowNum; rowSt++ {
		for col := 0; col < colNum; col++ {
			mymat[col] = true
		}
		for rowEd := rowSt; rowEd < rowNum; rowEd++ {
			length := 0
			for col := 0; col < colNum; col++ {
				if mat[rowEd][col] == 0 {
					mymat[col] = false
				}
				if mymat[col] {
					length++
					count += length
				} else {
					length = 0
				}
			}
		}
	}
	return count
}

// 28 ms, faster than 79.41%
// 6.1 MB, less than 100.00%
func numSubmat3(mat [][]int) int {
	count := 0
	rowNum := len(mat)
	colNum := len(mat[0])

	heights := make([]int, colNum)
	dpCount := make([]int, colNum)
	hStack := newStack()
	for row := 0; row < rowNum; row++ {
		hStack.clear()

		for col := 0; col < colNum; col++ {
			if mat[row][col] == 1 {
				heights[col]++
			} else {
				heights[col] = 0
			}

			for hStack.size() > 0 && heights[hStack.top()] >= heights[col] {
				hStack.pop()
			}

			if hStack.size() == 0 {
				dpCount[col] = heights[col] * (col + 1)
			} else {
				preCol := hStack.top()
				dpCount[col] = dpCount[preCol] + heights[col]*(col-preCol)
			}

			hStack.push(col)
		}

		for col := 0; col < colNum; col++ {
			count += dpCount[col]
		}
	}
	return count
}

type stack struct {
	data []int
}

func newStack() *stack      { return &stack{data: make([]int, 0)} }
func (s *stack) size() int  { return len(s.data) }
func (s *stack) push(v int) { s.data = append(s.data, v) }
func (s *stack) top() int   { return s.data[len(s.data)-1] }
func (s *stack) pop()       { s.data = s.data[:len(s.data)-1] }
func (s *stack) clear()     { s.data = s.data[:0] }
