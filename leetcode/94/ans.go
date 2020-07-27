package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0ms, 100%
// 2.1MB
func inorderTraversal(root *TreeNode) (ans []int) {
	if root == nil {
		return
	}

	ans = append(ans, inorderTraversal(root.Left)...)
	ans = append(ans, root.Val)
	ans = append(ans, inorderTraversal(root.Right)...)
	return
}

// 0ms, 100%
// 2MB, 100%
func inorderTraversalIter(root *TreeNode) (ans []int) {
	stack := make(nodeStack, 0)
	for root != nil {
		// go left-most
		for root.Left != nil {
			stack.push(root)
			root = root.Left
		}
		ans = append(ans, root.Val)
		// try right, or parent
		for root.Right == nil && len(stack) > 0 {
			root = stack.pop()
			ans = append(ans, root.Val)
		}
		root = root.Right
	}
	return
}

type nodeStack []*TreeNode

func (s *nodeStack) push(node *TreeNode) {
	(*s) = append((*s), node)
}

func (s *nodeStack) pop() *TreeNode {
	old := *s
	newSize := len(old) - 1
	(*s) = old[:newSize]
	return old[newSize]
}
