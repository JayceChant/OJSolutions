package main

import "strconv"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0ms
// 2.3MB
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}
	str := strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		return []string{str}
	}

	ans = nil
	if root.Left != nil {
		btp(root.Left, str)
	}

	if root.Right != nil {
		btp(root.Right, str)
	}

	return ans
}

var (
	ans []string
)

func btp(root *TreeNode, str string) {
	str = str + "->" + strconv.Itoa(root.Val)
	if root.Left == nil && root.Right == nil {
		ans = append(ans, str)
	}

	if root.Left != nil {
		btp(root.Left, str)
	}

	if root.Right != nil {
		btp(root.Right, str)
	}
}
