package main

import "container/list"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 8ms
// 5.2MB
func minDepthList(root *TreeNode) int {
	if root == nil {
		return 0
	}
	l := list.New()
	root.Val = 1
	l.PushBack(root)
	for l.Len() > 0 {
		e := l.Front()
		l.Remove(e)
		node := e.Value.(*TreeNode)
		if node.Left == nil && node.Right == nil {
			return node.Val
		}

		if node.Left != nil {
			node.Left.Val = node.Val + 1
			l.PushBack(node.Left)
		}

		if node.Right != nil {
			node.Right.Val = node.Val + 1
			l.PushBack(node.Right)
		}
	}
	return 0
}

// 0ms
// 5.1MB
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	depth := 1
	layer := []*TreeNode{root}
	nextLayer := make([]*TreeNode, 0)
	for {
		for _, node := range layer {
			if node.Left == nil && node.Right == nil {
				return depth
			}

			if node.Left != nil {
				nextLayer = append(nextLayer, node.Left)
			}

			if node.Right != nil {
				nextLayer = append(nextLayer, node.Right)
			}
		}
		layer, nextLayer = nextLayer, layer[:0]
		depth++
	}
}
