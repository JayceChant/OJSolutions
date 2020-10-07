package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 12ms (min 0ms)
// 5.7MB
func isBalanced(root *TreeNode) bool {
	return root == nil || balance(root) >= 0
}

func balance(root *TreeNode) int {
	l := 0
	if root.Left != nil {
		l = balance(root.Left)
		if l < 0 {
			return -1
		}
	}

	r := 0
	if root.Right != nil {
		r = balance(root.Right)
		if r < 0 {
			return -1
		}
	}
	if l < r {
		l, r = r, l
	}
	if l-r > 1 {
		return -1
	}
	return l + 1
}
