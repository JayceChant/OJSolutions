package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// no opt
// 20ms
// 3.3MB
// last node comparasion
// 4ms
// 3.3MB
func insertionSortList(head *ListNode) *ListNode {
	sorted := new(ListNode)
	var prev, node, next *ListNode
	for head != nil {
		if node != nil && node.Val < head.Val {
			// last node optimization
			prev = node
		} else {
			prev = sorted
		}
		node = head
		head = node.Next
		next = prev.Next
		// find insert point
		for next != nil && next.Val < node.Val {
			prev = next
			next = next.Next
		}
		prev.Next = node
		node.Next = next
	}
	return sorted.Next
}
