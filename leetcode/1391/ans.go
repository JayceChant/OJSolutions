package main

// 188ms
// 7.2MB
func hasValidPath(grid [][]int) bool {
	m := len(grid)
	n := len(grid[0])
	if m == 1 && n == 1 {
		return true
	}

	if grid[0][0] == 5 {
		return false
	}

	uf := make(unionFind, m*n)
	uf.init()
	for r := 0; r < m; r++ {
		for c := 0; c < n; c++ {
			gtype := grid[r][c] - 1
			id := r*n + c
			// from left
			if c > 0 && validFrom[gtype][0][grid[r][c-1]-1] {
				uf.union(id-1, id)
			}

			// from up
			if r > 0 && validFrom[gtype][1][grid[r-1][c]-1] {
				uf.union(id-n, id)
			}
		}
	}
	return uf.isUnited(0, m*n-1)
}

type unionFind []int

func (uf unionFind) init() {
	for i := 0; i < len(uf); i++ {
		uf[i] = i
	}
}

func (uf unionFind) findRoot(n int) int {
	root := uf[n]
	if root == n {
		return n
	}

	root = uf.findRoot(root)
	uf[n] = root
	return root
}

func (uf unionFind) union(a, b int) {
	ra := uf.findRoot(a)
	rb := uf.findRoot(b)
	uf[rb] = ra
}

func (uf unionFind) isUnited(a, b int) bool {
	return uf.findRoot(a) == uf.findRoot(b)
}

var (
	validFrom = [6][2][6]bool{
		{ // -
			{true, false, false, true, false, true},    // left
			{false, false, false, false, false, false}, // up
		},
		{ // |
			{false, false, false, false, false, false}, // left
			{false, true, true, true, false, false},    // up
		},
		{ // ┓
			{true, false, false, true, false, true},    // left
			{false, false, false, false, false, false}, // up
		},
		{ // ┏
			{false, false, false, false, false, false}, // left
			{false, false, false, false, false, false}, // up
		},
		{ // ┛
			{true, false, false, true, false, true}, // left
			{false, true, true, true, false, false}, // up
		},
		{ // ┗
			{false, false, false, false, false, false}, // left
			{false, true, true, true, false, false},    // up
		},
	}
)
