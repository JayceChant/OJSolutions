package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

const (
	maxRang = 20000
)

func removeDuplicateNodes(head *ListNode) *ListNode {
	var existed [maxRang + 1]bool
	ptr := head
	for ptr != nil {
		existed[ptr.Val] = true
		for ptr.Next != nil && existed[ptr.Next.Val] {
			ptr.Next = ptr.Next.Next
			continue
		}
		ptr = ptr.Next
	}
	return head
}
