package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

// use preHead as sentinel
// 4ms, 61.40%
// 2.5MB, 63.64%

// 0ms, 100%
// 2524, 72.73%
func mergeTwoListsMerge(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, tail *ListNode
	if l1.Val < l2.Val {
		head = l1
		tail = head
		l1 = l1.Next
	} else {
		head = l2
		tail = head
		l2 = l2.Next
	}

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next = l1
			tail = tail.Next
			l1 = l1.Next
		} else {
			tail.Next = l2
			tail = tail.Next
			l2 = l2.Next
		}
	}

	if l1 != nil {
		tail.Next = l1
	} else {
		tail.Next = l2
	}

	return head
}

// Insert
// 4ms
// 2520, 100%
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	if l1.Val > l2.Val {
		l1, l2 = l2, l1
	}
	cur := l1
	for cur.Next != nil && l2 != nil {
		for cur.Next != nil && cur.Next.Val <= l2.Val {
			cur = cur.Next
		}
		// cur <= l2 < cur.Next, swap l2 and cur.Next
		cur.Next, l2 = l2, cur.Next
		cur = cur.Next
	}

	if l2 != nil {
		cur.Next = l2
	}

	return l1
}
