package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// Recursive
func reverseListRec(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	head, tail := reverseRec(head)
	tail.Next = nil
	return head
}

func reverseRec(head *ListNode) (*ListNode, *ListNode) {
	if head.Next == nil {
		return head, head
	}

	newHead, tail := reverseRec(head.Next)
	tail.Next = head
	return newHead, head
}

// Iteration
// double 100%
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	last := (*ListNode)(nil)
	next := head.Next
	head.Next = last
	for next != nil {
		last, head, next = head, next, next.Next
		head.Next = last
	}
	return head
}
