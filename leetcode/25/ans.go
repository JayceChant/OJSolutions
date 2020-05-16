package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	if k == 1 {
		return head
	}
	cur := head
	var gHead, gTail, tail *ListNode
	for {
		gHead = cur
		if gHead == nil {
			return head
		}
		for i := 0; i < k; i++ {
			if cur == nil {
				tail.Next = gHead
				return head
			}
			gTail = cur
			cur = cur.Next
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

func reverseSubGroup(prev *ListNode) *ListNode {
	cur := prev.Next
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
