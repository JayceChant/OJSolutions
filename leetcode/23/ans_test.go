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
		// 为了代码简洁，不处理多余的 '->'
		sb.WriteString(fmt.Sprintf("%d->", l.Val))
		l = l.Next
	}
	return sb.String()
}

func listsFromSlices(aa [][]int) []*ListNode {
	ll := make([]*ListNode, 0, len(aa))
	head := new(ListNode)
	for _, a := range aa {
		head.Next = nil // 避免空数组没有覆盖
		tail := head
		for _, v := range a {
			tail.Next = &ListNode{Val: v}
			tail = tail.Next
		}
		ll = append(ll, head.Next)
	}
	return ll
}

var tests = []struct {
	name  string
	cases [][]int
	want  string
}{
	{
		name:  "0",
		cases: [][]int{},
		want:  "",
	},
	{
		name: "00",
		cases: [][]int{
			{},
		},
		want: "",
	},
	{
		name: "1",
		cases: [][]int{
			{1, 4, 5},
			{1, 3, 4},
			{2, 6},
		},
		want: "1->1->2->3->4->4->5->6->",
	},
	{
		name: "2",
		cases: [][]int{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5, 5},
			{2, 3, 4},
			{2, 6},
			{9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9}, {9},
		},
		want: "1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->2->2->2->2->2->2->2->2->2->2->2->2->2->2->2->2->2->2->2->2->2->2->3->3->3->3->3->3->3->3->3->3->3->3->3->3->3->3->3->3->3->3->3->4->4->4->4->4->4->4->4->4->4->4->4->4->4->4->4->4->4->4->4->4->5->5->5->5->5->5->5->5->5->5->5->5->5->5->5->5->5->5->5->5->6->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->9->",
	},
}

func Test_mergeKLists(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// 链表每次要重新生成
			if got := mergeKListsOneByOne(listsFromSlices(tt.cases)); got.String() != tt.want {
				t.Errorf("mergeKListsOneByOne() = %s, want %s", got, tt.want)
			}
			if got := mergeKListsInsert(listsFromSlices(tt.cases)); got.String() != tt.want {
				t.Errorf("mergeKListsInsert() = %s, want %s", got, tt.want)
			}
			if got := mergeKListsInsertDAC(listsFromSlices(tt.cases)); got.String() != tt.want {
				t.Errorf("mergeKListsInsertDAC() = %s, want %s", got, tt.want)
			}
			if got := mergeKListsWithHeap(listsFromSlices(tt.cases)); got.String() != tt.want {
				t.Errorf("mergeKListsWithHeap() = %s, want %s", got, tt.want)
			}
			if got := mergeKListsWithHeap2(listsFromSlices(tt.cases)); got.String() != tt.want {
				t.Errorf("mergeKListsWithHeap2() = %s, want %s", got, tt.want)
			}
			if got := mergeKListsWithHeap3(listsFromSlices(tt.cases)); got.String() != tt.want {
				t.Errorf("mergeKListsWithHeap3() = %s, want %s", got, tt.want)
			}
			if got := mergeKListsWithBuiltinHeap(listsFromSlices(tt.cases)); got.String() != tt.want {
				t.Errorf("mergeKListsWithBuiltinHeap() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Benchmark_mergeKListsOneByOne(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mergeKListsOneByOne(listsFromSlices(tt.cases))
		}
	}
}

func Benchmark_mergeKListsInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mergeKListsInsert(listsFromSlices(tt.cases))
		}
	}
}

func Benchmark_mergeKListsInsertDAC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mergeKListsInsertDAC(listsFromSlices(tt.cases))
		}
	}
}

func Benchmark_mergeKListsWithHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mergeKListsWithHeap(listsFromSlices(tt.cases))
		}
	}
}

func Benchmark_mergeKListsWithHeap2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mergeKListsWithHeap2(listsFromSlices(tt.cases))
		}
	}
}

func Benchmark_mergeKListsWithHeap3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mergeKListsWithHeap3(listsFromSlices(tt.cases))
		}
	}
}

func Benchmark_mergeKListsWithBuiltinHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mergeKListsWithBuiltinHeap(listsFromSlices(tt.cases))
		}
	}
}
