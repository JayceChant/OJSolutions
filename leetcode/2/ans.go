package main

// ListNode ...
type ListNode struct {
	Val  int
	Next *ListNode
}

var head = new(ListNode)

// 8 ms, faster than 89.39%
// 5 MB, less than 65.71%
// reuse l1:
// 12 ms, faster than 63.85%
// 4.8 MB, less than 92.40%
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	tail := head
	tmp := 0
	for l1 != nil && l2 != nil {
		tmp = l1.Val + l2.Val + tmp
		node := l1
		node.Val = tmp % 10
		tail.Next = node
		tmp /= 10
		l1 = l1.Next
		l2 = l2.Next
		tail = node
	}

	if l2 != nil {
		l1 = l2
	}

	for l1 != nil {
		tmp = l1.Val + tmp
		node := l1
		node.Val = tmp % 10
		tail.Next = node
		tmp /= 10
		l1 = l1.Next
		tail = node
	}

	if tmp != 0 {
		node := new(ListNode)
		node.Val = tmp % 10
		tail.Next = node
	}

	return head.Next
}

func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	tail := head
	tmp := 0
	for l1 != nil || l2 != nil || tmp > 0 {
		if l1 != nil {
			tmp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmp += l2.Val
			l2 = l2.Next
		}
		tail.Next = &ListNode{Val: tmp % 10}
		tmp /= 10
		tail = tail.Next
	}
	return head.Next
}
