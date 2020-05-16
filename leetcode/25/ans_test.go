package main

import (
	"fmt"
	"strings"
	"testing"
)

var _ fmt.Stringer = (*ListNode)(nil)

// String 方法方便调试输出
func (head *ListNode) String() string {
	b := strings.Builder{}
	for ; head != nil; head = head.Next {
		// 为了省事，不特殊处理最后的逗号
		b.WriteString(fmt.Sprintf("%d%s", head.Val, ","))
	}
	return b.String()
}

func sliceToList(s []int) *ListNode {
	// 为了不用特殊处理，头结点留空
	head := new(ListNode)
	last := head
	var cur *ListNode
	for _, v := range s {
		cur = new(ListNode)
		cur.Val = v
		last.Next = cur
		last = cur
	}
	// 跳过头结点
	return head.Next
}

var tests = []struct {
	name string
	k    int
	s    []int
	want string
}{
	{
		name: "1",
		k:    1,
		s:    []int{1, 2, 3, 4, 5, 6, 7},
		want: "1,2,3,4,5,6,7,",
	},
	{
		name: "2",
		k:    2,
		s:    []int{1, 2, 3, 4, 5, 6, 7},
		want: "2,1,4,3,6,5,7,",
	},
	{
		name: "3",
		k:    3,
		s:    []int{1, 2, 3, 4, 5, 6, 7, 8},
		want: "3,2,1,6,5,4,7,8,",
	},
	{
		name: "39",
		k:    3,
		s:    []int{1, 2, 3, 4, 5, 6, 7, 8, 9},
		want: "3,2,1,6,5,4,9,8,7,",
	},
}

func TestReverseKGroup(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := sliceToList(tt.s)
			// t.Log(head)
			got := reverseKGroup(head, tt.k).String()
			if got != tt.want {
				t.Errorf("reverseKGroup() = %v; want %v", got, tt.want)
			}
		})
	}
}
