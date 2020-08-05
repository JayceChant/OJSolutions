package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// house robber III
// DP
// 8ms (min 0ms)
// 5.2MB
func rob(root *TreeNode) int {
	act, skip := robNode(root)
	return max(act, skip)
}

func robNode(root *TreeNode) (act int, skip int) {
	if root == nil {
		return 0, 0
	}

	lact, lskip := robNode(root.Left)
	ract, rskip := robNode(root.Right)
	return lskip + rskip + root.Val, max(lact, lskip) + max(ract, rskip)
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}
