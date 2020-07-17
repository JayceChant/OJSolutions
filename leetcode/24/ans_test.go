package main

import (
	"testing"
)

var tests = []struct {
	name  string
	array []int
}{
	{
		name:  "1",
		array: []int{1, 2, 3},
	},
}

func Test_swapPairs(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			list := linkedListFromArray(tt.array)
			isValid(t, answerFromLinkedList(list), swapPairs(list))
		})
	}
}

var (
	dummy = new(ListNode)
)

func linkedListFromArray(array []int) *ListNode {
	dummy.Next = nil
	tail := dummy
	for _, num := range array {
		tail.Next = &ListNode{
			Val: num,
		}
		tail = tail.Next
	}
	return dummy.Next
}

func answerFromLinkedList(head *ListNode) []*ListNode {
	ans := make([]*ListNode, 0)
	for head != nil && head.Next != nil {
		ans = append(ans, head.Next, head)
		head = head.Next.Next
	}
	return ans
}

func isValid(t *testing.T, answer []*ListNode, head *ListNode) {
	for _, node := range answer {
		if head != node {
			t.Errorf("swapPairs() = %v, want %v", head, node)
			return
		}
		head = head.Next
	}
}
