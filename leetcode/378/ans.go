package main

func kthSmallest(matrix [][]int, k int) int {
	n := len(matrix)
	var offset int
	var e *node
	h := &heap{data: make([]*node, 0)}
	in := make([][]uint8, n)
	for i := 0; i < n; i++ {
		in[i] = make([]uint8, n)
	}
	if k <= n*n/2 {
		offset = 1
		e = &node{
			val: matrix[0][0],
			row: 0,
			col: 0,
		}
		h.min = true
		for i := 0; i < n; i++ {
			in[0][i] = 1
			in[i][0] = 1
		}
	} else {
		offset = -1
		e = &node{
			val: matrix[n-1][n-1],
			row: n - 1,
			col: n - 1,
		}
		h.min = false
		for i := 0; i < n; i++ {
			in[i][n-1] = 1
			in[n-1][i] = 1
		}
		k = n*n - k + 1
	}
	h.push(e)

	for i := 1; i < k; i++ {
		e = h.pop()
		if row := e.row + offset; row >= 0 && row < n {
			in[row][e.col]++
			if in[row][e.col] >= 2 {
				h.push(&node{
					val: matrix[row][e.col],
					row: row,
					col: e.col,
				})
			}
		}
		if col := e.col + offset; col >= 0 && col < n {
			in[e.row][col]++
			if in[e.row][col] >= 2 {
				h.push(&node{
					val: matrix[e.row][col],
					row: e.row,
					col: col,
				})
			}
		}
	}
	e = h.pop()
	return e.val
}

type node struct {
	val int
	row int
	col int
}

type heap struct {
	min  bool
	data []*node
}

func (h *heap) before(i, j int) bool {
	if h.min {
		return h.data[i].val < h.data[j].val
	}
	return h.data[i].val > h.data[j].val
}

func (h *heap) push(e *node) {
	h.data = append(h.data, e)
	child := len(h.data) - 1
	for child > 0 {
		parent := (child - 1) / 2
		if h.before(child, parent) {
			h.data[parent], h.data[child] = h.data[child], h.data[parent]
		}
		child = parent
	}
}

func (h *heap) pop() *node {
	tail := len(h.data) - 1
	ret := h.data[0]
	h.data[0] = h.data[tail]
	parent := 0
	for {
		child := parent*2 + 1
		if child >= tail {
			break
		}
		if child+1 < tail && h.before(child+1, child) {
			child++
		}
		if h.before(parent, child) {
			break
		}
		h.data[parent], h.data[child] = h.data[child], h.data[parent]
		parent = child
	}
	h.data = h.data[:tail]
	return ret
}

func kthSmallest2(matrix [][]int, k int) int {
	return 0
}
