package main

import (
	"testing"
)

func (l *ListNode) intVal() int {
	num := 0
	base := 1
	for l != nil {
		num += l.Val * base
		l = l.Next
		base *= 10
	}
	return num
}

func listFromInt(num int) *ListNode {
	head := new(ListNode)
	tail := head
	for num != 0 {
		tail.Next = &ListNode{Val: num % 10}
		tail = tail.Next
		num /= 10
	}
	return head.Next
}

var tests = []struct {
	name string
	num1 int
	num2 int
	want int
}{
	{
		name: "1",
		num1: 342,
		num2: 465,
		want: 807,
	},
	{
		name: "2",
		num1: 999999999542,
		num2: 465,
		want: 1000000000007,
	},
	{
		name: "3",
		num1: 342,
		num2: 9999999999665,
		want: 10000000000007,
	},
}

func Test_addTwoNumbers(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addTwoNumbers(listFromInt(tt.num1), listFromInt(tt.num2)).intVal(); got != tt.want {
				t.Errorf("addTwoNumbers() = %v, want %v", got, tt.want)
			}

			if got := addTwoNumbers2(listFromInt(tt.num1), listFromInt(tt.num2)).intVal(); got != tt.want {
				t.Errorf("addTwoNumbers2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_addTwoNumbers(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			addTwoNumbers(listFromInt(tt.num1), listFromInt(tt.num2))
		}
	}
}

func Benchmark_addTwoNumbers2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			addTwoNumbers2(listFromInt(tt.num1), listFromInt(tt.num2))
		}
	}
}
