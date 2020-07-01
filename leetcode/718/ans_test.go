package main

import "testing"

var tests = []struct {
	name string
	A    []int
	B    []int
	want int
}{
	{
		name: "0",
		A:    []int{1},
		B:    []int{3},
		want: 0,
	},
	{
		name: "1",
		A:    []int{1, 2, 3, 2, 1},
		B:    []int{3, 2, 1, 4, 7},
		want: 3,
	},
	{
		name: "fail",
		A:    []int{1, 0, 0, 0, 1, 0, 0, 1, 0, 0},
		B:    []int{0, 1, 1, 1, 0, 1, 1, 1, 0, 0},
		want: 3,
	},
}

func Test_findLength(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// copy before call the function that has side effect
			// cA := make([]int, len(tt.A))
			// copy(cA, tt.A)
			// cB := make([]int, len(tt.B))
			// copy(cB, tt.B)
			if got := findLengthDP(tt.A, tt.B); got != tt.want {
				t.Errorf("findLengthDP() = %v, want %v", got, tt.want)
			}
			if got := findLengthOffset(tt.A, tt.B); got != tt.want {
				t.Errorf("findLengthOffset() = %v, want %v", got, tt.want)
			}
			if got := findLengthHash(tt.A, tt.B); got != tt.want {
				t.Errorf("findLengthHash() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_findLengthDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			findLengthDP(tt.A, tt.B)
		}
	}
}

func Benchmark_findLengthOffset(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			findLengthOffset(tt.A, tt.B)
		}
	}
}

func Benchmark_findLengthHash(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			findLengthHash(tt.A, tt.B)
		}
	}
}
