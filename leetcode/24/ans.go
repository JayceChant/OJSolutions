package main

// ListNode ..
type ListNode struct {
	Val  int
	Next *ListNode
}

// 0ms, 100%
// 2.1MB, 20%
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	dummy := new(ListNode)
	dummy.Next = head

	left := dummy
	right := head.Next
	for right != nil {
		left.Next, right.Next, head.Next = right, head, right.Next
		if head.Next == nil || head.Next.Next == nil {
			break
		}
		left = head
		head = left.Next
		right = head.Next
	}

	return dummy.Next
}
