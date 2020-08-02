package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0ms
// 2.9MB
func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatNode(root, nil)
}

func flatNode(root *TreeNode, rightRoot *TreeNode) {
	if root.Right == nil {
		root.Right = rightRoot
	} else {
		flatNode(root.Right, rightRoot)
	}

	if root.Left == nil {
		return
	}

	flatNode(root.Left, root.Right)
	root.Right = root.Left
	root.Left = nil
}
