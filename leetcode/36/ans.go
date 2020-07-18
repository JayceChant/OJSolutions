package main

// 4ms, 67.10%
// 3.1MB, 9.52%
func isValidSudokuBool(board [][]byte) bool {
	rows := make([][]bool, 9)
	cols := make([][]bool, 9)
	grids := make([][]bool, 9)

	for i := 0; i < 9; i++ {
		rows[i] = make([]bool, 9)
		cols[i] = make([]bool, 9)
		grids[i] = make([]bool, 9)
	}

	for r := range board {
		for c, num := range board[r] {
			if num == '.' {
				continue
			}

			num -= '1'
			if rows[r][num] || cols[c][num] || grids[r-r%3+c/3][num] {
				// fmt.Println(r, c)
				return false
			}

			rows[r][num] = true
			cols[c][num] = true
			grids[r-r%3+c/3][num] = true

		}
	}
	return true
}

// 4ms, 67.10%
// 2.8MB, 57.14%
func isValidSudokuBit(board [][]byte) bool {
	var row uint16
	cols := make([]uint16, 9)
	grids := make([]uint16, 3)

	for r := range board {
		row = 0
		if r%3 == 0 {
			grids[0] = 0
			grids[1] = 0
			grids[2] = 0
		}
		for c, num := range board[r] {
			if num == '.' {
				continue
			}

			mask := masks[num-'1']

			if row&mask > 0 || cols[c]&mask > 0 || grids[c/3]&mask > 0 {
				// fmt.Println(r, c)
				return false
			}

			row |= mask
			cols[c] |= mask
			grids[c/3] |= mask
		}
	}
	return true
}

var (
	masks = []uint16{
		1,
		1 << 1,
		1 << 2,
		1 << 3,
		1 << 4,
		1 << 5,
		1 << 6,
		1 << 7,
		1 << 8,
	}
)
