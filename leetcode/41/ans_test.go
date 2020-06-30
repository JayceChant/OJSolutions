package main

import "testing"

var tests = []struct {
	name string
	nums []int
	want int
}{
	{
		name: "0",
		nums: []int{},
		want: 1,
	},
	{
		name: "1",
		nums: []int{1, 2, 0},
		want: 3,
	},
	{
		name: "2",
		nums: []int{3, 4, -1, 1},
		want: 2,
	},
	{
		name: "3",
		nums: []int{7, 8, 9, 11, 12},
		want: 1,
	},
	{
		name: "fail",
		nums: []int{1, 2},
		want: 3,
	},
}

func Test_firstMissingPositive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := firstMissingPositive(tt.nums); got != tt.want {
			// 	t.Errorf("firstMissingPositive() = %v, want %v", got, tt.want)
			// }

			if got := firstMissingPositive2(tt.nums); got != tt.want {
				t.Errorf("firstMissingPositive2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_firstMissingPositive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			args := make([]int, len(tt.nums))
			copy(args, tt.nums)
			firstMissingPositive(args)
		}
	}
}

func Benchmark_firstMissingPositive2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			args := make([]int, len(tt.nums))
			copy(args, tt.nums)
			firstMissingPositive2(args)
		}
	}
}
