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
	{
		name: "3",
		cases: [][]int{
			{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			{2}, {3}, {4}, {5}, {6}, {7}, {8}, {9}, {10}, {11}, {12}, {13}, {14}, {15}, {16}, {17}, {18}, {19}, {20}, {21}, {22}, {23}, {24}, {25}, {26}, {27}, {28}, {29}, {30}, {31}, {32}, {33}, {34}, {35}, {36}, {37}, {38}, {39}, {40}, {41}, {42}, {43}, {44}, {45}, {46}, {47}, {48}, {49}, {50}, {51}, {52}, {53}, {54}, {55}, {56}, {57}, {58}, {59}, {60}, {61}, {62}, {63}, {64}, {65}, {66}, {67}, {68}, {69}, {70}, {71}, {72}, {73}, {74}, {75}, {76}, {77}, {78}, {79}, {80}, {81}, {82}, {83}, {84}, {85}, {86}, {87}, {88}, {89}, {90}, {91}, {92}, {93}, {94}, {95}, {96}, {97}, {98}, {99}, {100},
		},
		want: "1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->1->2->3->4->5->6->7->8->9->10->11->12->13->14->15->16->17->18->19->20->21->22->23->24->25->26->27->28->29->30->31->32->33->34->35->36->37->38->39->40->41->42->43->44->45->46->47->48->49->50->51->52->53->54->55->56->57->58->59->60->61->62->63->64->65->66->67->68->69->70->71->72->73->74->75->76->77->78->79->80->81->82->83->84->85->86->87->88->89->90->91->92->93->94->95->96->97->98->99->100->",
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
			if got := mergeKListsWithNHeap(listsFromSlices(tt.cases)); got.String() != tt.want {
				t.Errorf("mergeKListsWithHeap() = %s, want %s", got, tt.want)
			}
			if got := mergeKListsWithHeapPushPop(listsFromSlices(tt.cases)); got.String() != tt.want {
				t.Errorf("mergeKListsWithHeap2() = %s, want %s", got, tt.want)
			}
			if got := mergeKListsWithHeapUpdateRemove(listsFromSlices(tt.cases)); got.String() != tt.want {
				t.Errorf("mergeKListsWithHeap3() = %s, want %s", got, tt.want)
			}
			if got := mergeKListsWithBuiltinHeap(listsFromSlices(tt.cases)); got.String() != tt.want {
				t.Errorf("mergeKListsWithBuiltinHeap() = %s, want %s", got, tt.want)
			}
		})
	}
}

func Benchmark_listsFromSlices(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			listsFromSlices(tt.cases)
		}
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

func Benchmark_mergeKListsWithNHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mergeKListsWithNHeap(listsFromSlices(tt.cases))
		}
	}
}

func Benchmark_mergeKListsWithHeapPushPop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mergeKListsWithHeapPushPop(listsFromSlices(tt.cases))
		}
	}
}

func Benchmark_mergeKListsWithHeapUpdateRemove(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			mergeKListsWithHeapUpdateRemove(listsFromSlices(tt.cases))
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
