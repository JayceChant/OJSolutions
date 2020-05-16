package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	// return reverseSubGroup(head)
	nextG := head
	var gHead, gTail, tail *ListNode
	for {
		gHead = nextG
		if gHead == nil {
			return head
		}
		for i := 0; i < k; i++ {
			if nextG == nil {
				tail.Next = gHead
				return head
			}
			gTail = nextG
			nextG = nextG.Next
		}
		gTail.Next = nil
		if tail == nil {
			head = reverseSubGroup(gHead)
		} else {
			tail.Next = reverseSubGroup(gHead)
		}
		tail = gHead
	}
}

func reverseSubGroup(head *ListNode) *ListNode {
	prev := head
	cur := prev.Next
	if cur == nil {
		return head
	}

	next := cur.Next
	prev.Next = nil

	for {
		cur.Next = prev
		if next == nil {
			return cur
		}
		prev = cur
		cur = next
		next = next.Next
	}
}
