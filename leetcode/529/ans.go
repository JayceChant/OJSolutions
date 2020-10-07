package main

// 28ms (24ms)
// 6.2MB
func updateBoard(board [][]byte, click []int) [][]byte {
	if board[click[0]][click[1]] == 'M' {
		board[click[0]][click[1]] = 'X'
		return board
	}

	bb = board
	rowNum = len(board)
	colNum = len(board[0])

	clickE(click[0], click[1], all)
	return board
}

var (
	bb             [][]byte
	rowNum, colNum int
)

func clickE(r, c int, from direction) {
	hasRight := c < colNum-1
	hasUp := r > 0
	hasLeft := c > 0
	hasDown := r < rowNum-1

	var count byte

	if hasRight && from&left != 0 && bb[r][c+1] == 'M' {
		count++
	}

	if hasRight && hasUp && from&downLeft != 0 && bb[r-1][c+1] == 'M' {
		count++
	}

	if hasUp && from&down != 0 && bb[r-1][c] == 'M' {
		count++
	}

	if hasUp && hasLeft && from&downRight != 0 && bb[r-1][c-1] == 'M' {
		count++
	}

	if hasLeft && from&right != 0 && bb[r][c-1] == 'M' {
		count++
	}

	if hasLeft && hasDown && from&upRight != 0 && bb[r+1][c-1] == 'M' {
		count++
	}

	if hasDown && from&up != 0 && bb[r+1][c] == 'M' {
		count++
	}

	if hasDown && hasRight && from&upLeft != 0 && bb[r+1][c+1] == 'M' {
		count++
	}

	if count > 0 {
		bb[r][c] = '0' + count
		return
	}

	bb[r][c] = 'B'

	if hasRight && from&left != 0 && bb[r][c+1] == 'E' {
		clickE(r, c+1, left)
	}

	if hasRight && hasUp && from&downLeft != 0 && bb[r-1][c+1] == 'E' {
		clickE(r-1, c+1, downLeft)
	}

	if hasUp && from&down != 0 && bb[r-1][c] == 'E' {
		clickE(r-1, c, down)
	}

	if hasUp && hasLeft && from&downRight != 0 && bb[r-1][c-1] == 'E' {
		clickE(r-1, c-1, downRight)
	}

	if hasLeft && from&right != 0 && bb[r][c-1] == 'E' {
		clickE(r, c-1, right)
	}

	if hasLeft && hasDown && from&upRight != 0 && bb[r+1][c-1] == 'E' {
		clickE(r+1, c-1, upRight)
	}

	if hasDown && from&up != 0 && bb[r+1][c] == 'E' {
		clickE(r+1, c, up)
	}

	if hasDown && hasRight && from&upLeft != 0 && bb[r+1][c+1] == 'E' {
		clickE(r+1, c+1, upLeft)
	}
}

type direction uint8

const (
	right     direction = 1
	up        direction = 2
	left      direction = 4
	down      direction = 8
	upRight             = up | right
	upLeft              = up | left
	downLeft            = down | left
	downRight           = down | right
	all                 = right | up | left | down
)
