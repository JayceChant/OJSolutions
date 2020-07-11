package main

func countSmallerBF(nums []int) []int {
	size := len(nums)
	counts := make([]int, size)
	for i := size - 1; i >= 0; i-- {
		for j := i + 1; j < size; j++ {
			if nums[j] < nums[i] {
				counts[i]++
			}
		}
	}
	return counts
}

// 8 ms, faster than 99.10% (min 4ms)
// 5.4 MB, less than 20.59% (min 4128kB)
func countSmallerMerge(nums []int) []int {
	size := len(nums)
	nodes := make([]*node, size)
	sorted := make([]*node, size)
	for i := 0; i < size; i++ {
		nodes[i] = &node{
			value: nums[i],
			index: i,
		}
	}

	step := 1
	width := 2
	for step < size {
		start := 0
		for start+width <= size {
			merge(nodes, sorted, start, start+step, start+width)
			start += width
		}
		if start < size {
			merge(nodes, sorted, start, start+step, size)
		}
		step = width
		width *= 2
		nodes, sorted = sorted, nodes
	}
	//fmt.Println(nodes)

	counts := make([]int, size)
	for i := 0; i < size; i++ {
		counts[nodes[i].index] = nodes[i].count
	}
	return counts
}

// merge [left, right) and [right, end) in from into sorted
func merge(from, to []*node, leftSt, rightSt, rightEd int) {
	toIdx := leftSt
	left := leftSt
	right := rightSt
	for toIdx < rightEd {
		if left == rightSt || (right < rightEd && from[left].value > from[right].value) {
			to[toIdx] = from[right]
			right++
		} else {
			from[left].count += (toIdx - left)
			to[toIdx] = from[left]
			left++
		}
		toIdx++
	}
}

type node struct {
	value int
	index int
	count int
}

// func (n *node) String() string {
// 	return fmt.Sprintf("[%d,%d,%d]", n.value, n.index, n.count)
// }

// 12 ms, faster than 77.48%
// 4.9 MB, less than 50.00%
func countSmallerAVL(nums []int) []int {
	size := len(nums)
	counts := make([]int, size)
	if size == 0 {
		return counts
	}
	counts[size-1] = 0
	if size == 1 {
		return counts
	}

	tree := &avlNode{
		value: nums[size-1],
		count: 1,
	}
	for i := size - 2; i >= 0; i-- {
		//fmt.Println(tree)
		counts[i] = tree.add(nums[i])
	}
	return counts
}

type avlNode struct {
	value  int
	height int
	count  int
	left   *avlNode
	right  *avlNode
}

func (n *avlNode) add(value int) (countSmaller int) {
	// count 包括子树的值，所以经过就要自增
	n.count++
	if value == n.value {
		if n.left == nil {
			return 0
		}
		return n.left.count
	}

	if value < n.value {
		if n.left == nil {
			n.left = &avlNode{
				value: value,
				count: 1,
			}
			if n.height == 0 {
				n.height = 1
			}
			return 0
		}

		countSmaller = n.left.add(value)
		n.left = n.left.rebalance()
		return
	}

	// value > n.value
	if n.right == nil {
		n.right = &avlNode{
			value: value,
			count: 1,
		}
		if n.height == 0 {
			n.height = 1
		}
		return n.count - 1
	}
	countSmaller = n.right.add(value)
	countSmaller += (n.count - n.right.count)
	n.right = n.right.rebalance()
	return
}

func (n *avlNode) rebalance() *avlNode {
	f := balanceFactor(n)

	if f >= 2 {
		if balanceFactor(n.left) >= 1 { // LL
			return n.rotateRight()
		}
		// LR
		n.left = n.left.rotateLeft()
		return n.rotateRight()
	}

	if f <= -2 {
		if balanceFactor(n.right) <= -1 { // RR
			return n.rotateLeft()
		}
		// RL
		n.right = n.right.rotateRight()
		return n.rotateLeft()
	}
	return n
}

func (n *avlNode) rotateRight() *avlNode {
	root := n.left
	// update count
	midCount := 0
	if root.right != nil {
		midCount = root.right.count
	}
	root.count, n.count = n.count, n.count-root.count+midCount
	// update link
	root.right, n.left = n, root.right
	// update height
	n.updateHeight()
	if root.height <= n.height {
		root.height = n.height + 1
	}
	return root
}

func (n *avlNode) rotateLeft() *avlNode {
	root := n.right
	// update count
	midCount := 0
	if root.left != nil {
		midCount = root.left.count
	}
	root.count, n.count = n.count, n.count-root.count+midCount
	// update link
	n.right, root.left = root.left, n
	// update height
	n.updateHeight()
	if root.height <= n.height {
		root.height = n.height + 1
	}
	return root
}

func (n *avlNode) updateHeight() {
	n.height = 0
	if n.left != nil {
		n.height = n.left.height + 1
	}
	if n.right != nil && n.height <= n.right.height {
		n.height = n.right.height + 1
	}
}

func balanceFactor(n *avlNode) int {
	if n == nil {
		return 0
	}
	lHeight := -1
	if n.left != nil {
		lHeight = n.left.height
	}
	rHeight := -1
	if n.right != nil {
		rHeight = n.right.height
	}
	return lHeight - rHeight
}

// count when add
// func (n *avlNode) countSmaller(value int) int {
// 	//fmt.Printf("[%d,%d],", value, n.value)
// 	if value == n.value {
// 		if n.left == nil {
// 			return 0
// 		}
// 		return n.left.count
// 	}

// 	if value < n.value {
// 		if n.left == nil {
// 			return 0
// 		}
// 		return n.left.countSmaller(value)
// 	}

// 	// value > n.value
// 	if n.right == nil {
// 		return n.count
// 	}
// 	return n.count - n.right.count + n.right.countSmaller(value)
// }

// func (n *avlNode) String() string {
// 	return fmt.Sprintf("[%v,%d=%d,%v]", n.left, n.value, n.count, n.right)
// }
