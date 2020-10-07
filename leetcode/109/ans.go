package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 8ms
// 6.1MB
func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}

	length := 0
	tail := head
	for tail != nil {
		length++
		tail = tail.Next
	}

	current = head
	tree := buildTree(length)
	return tree
}

var (
	current *ListNode
)

func buildTree(length int) *TreeNode {
	root := new(TreeNode)

	llen := length / 2
	if llen > 0 {
		root.Left = buildTree(llen)
	}

	root.Val = current.Val
	current = current.Next

	length = length - llen - 1
	if length > 0 {
		root.Right = buildTree(length)
	}
	return root
}
