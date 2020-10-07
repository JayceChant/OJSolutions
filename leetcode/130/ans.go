package main

// 24ms
// 6MB
func solve(board [][]byte) {
	height = len(board)
	if height == 0 {
		return
	}

	width = len(board[0])
	if width == 0 {
		return
	}
	gb = board

	// 1st and last row
	last := height - 1
	for c := 0; c < width; c++ {
		if board[0][c] == 'O' {
			dfs(0, c)
		}

		if board[last][c] == 'O' {
			dfs(last, c)
		}
	}

	// 1st and last col
	last = width - 1
	for r := 1; r < height-1; r++ {
		if board[r][0] == 'O' {
			dfs(r, 0)
		}

		if board[r][last] == 'O' {
			dfs(r, last)
		}
	}

	for r := 0; r < height; r++ {
		for c := 0; c < width; c++ {
			switch board[r][c] {
			case 'O':
				board[r][c] = 'X'
			case 'Q':
				board[r][c] = 'O'
			}
		}
	}
}

var (
	width, height int
	gb            [][]byte
)

func dfs(r, c int) {
	gb[r][c] = 'Q'
	// up
	if r > 0 && gb[r-1][c] == 'O' {
		dfs(r-1, c)
	}
	// down
	if r < height-1 && gb[r+1][c] == 'O' {
		dfs(r+1, c)
	}
	// left
	if c > 0 && gb[r][c-1] == 'O' {
		dfs(r, c-1)
	}
	// right
	if c < width-1 && gb[r][c+1] == 'O' {
		dfs(r, c+1)
	}
}
