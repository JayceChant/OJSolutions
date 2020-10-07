package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 24ms (min 12ms)
// 6.8MB
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	if p.Right != nil {
		p = p.Right
		s := p.Left
		for s != nil {
			p = s
			s = s.Left
		}
		return p
	}
	// p.Right == nil
	var next *TreeNode
	for root != p {
		if p.Val < root.Val {
			next = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return next
}
