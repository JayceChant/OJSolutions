package main

import (
	"fmt"
	"strings"
	"testing"
)

var _ fmt.Stringer = (*ListNode)(nil)

func (l *ListNode) String() string {
	sb := strings.Builder{}
	for l != nil {
		// 为了代码简洁，不处理多余的 ','
		sb.WriteString(fmt.Sprintf("%d,", l.Val))
		l = l.Next
	}
	return sb.String()
}

func listFromSlice(a []int) *ListNode {
	head := new(ListNode)
	tail := head
	for _, v := range a {
		tail.Next = &ListNode{Val: v}
		tail = tail.Next
	}
	return head.Next
}

var tests = []struct {
	name string
	args []int
	want string
}{
	{
		name: "0",
		args: []int{},
		want: "",
	},
	{
		name: "1",
		args: []int{1, 2, 3, 3, 2, 1},
		want: "1,2,3,",
	},
	{
		name: "2",
		args: []int{1, 1, 1, 1, 2},
		want: "1,2,",
	},
}

func Test_removeDuplicateNodes(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeDuplicateNodes(listFromSlice(tt.args)).String(); got != tt.want {
				t.Errorf("removeDuplicateNodes() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Benchmark_removeDuplicateNodes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			removeDuplicateNodes(listFromSlice(tt.args))
		}
	}
}
