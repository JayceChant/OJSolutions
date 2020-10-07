package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distributeCoins(root *TreeNode) int {
	if root == nil {
		return 0
	}
	step = 0
	diff(root)
	return step
}

var (
	step int
)

func diff(root *TreeNode) (surplus int) {
	surplus = root.Val - 1

	if root.Left != nil {
		surplus += diff(root.Left)
	}

	if root.Right != nil {
		surplus += diff(root.Right)
	}

	step += abs(surplus)
	return
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}
