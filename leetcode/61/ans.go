package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// 0ms
// 2.5MB
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	length := 1
	tail := head
	for tail.Next != nil {
		tail = tail.Next
		length++
	}

	k %= length
	if k == 0 {
		return head
	}

	tail.Next = head
	k = length - k
	for i := 0; i < k; i++ {
		tail = tail.Next
	}
	head = tail.Next
	tail.Next = nil
	return head
}
