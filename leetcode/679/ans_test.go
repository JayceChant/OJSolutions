package main

import "testing"

var tests = []struct {
	name string
	nums []int
	want bool
}{
	{
		name: "fail true",
		nums: []int{1, 5, 9, 1},
		want: false,
	},
	{
		name: "fail false",
		nums: []int{1, 3, 4, 6},
		want: true,
	},
	{
		name: "fail false2",
		nums: []int{3, 3, 8, 8},
		want: true,
	},
}

func Test_judgePoint24(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bruteForce(tt.nums); got != tt.want {
				t.Errorf("bruteForce() = %v, want %v", got, tt.want)
			}
		})
		t.Run(tt.name, func(t *testing.T) {
			if got := judgePoint24(tt.nums); got != tt.want {
				t.Errorf("judgePoint24() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_bruteForce(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			bruteForce(tt.nums)
		}
	}
}

func Benchmark_judgePoint24(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			judgePoint24(tt.nums)
		}
	}
}
