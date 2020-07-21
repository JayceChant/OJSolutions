package main

import "testing"

func Test_findMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "edge0",
			nums: []int{3, 3, 1, 3},
			want: 1,
		},
		{
			name: "edge1",
			nums: []int{1},
			want: 1,
		},
		{
			name: "edge2",
			nums: []int{2, 2},
			want: 2,
		},
		{
			name: "edge3",
			nums: []int{3, 3, 3},
			want: 3,
		},
		{
			name: "edge8",
			nums: []int{8, 8, 8, 8, 8, 8, 8, 8},
			want: 8,
		},
		{
			name: "fail",
			nums: []int{1, 3, 3},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
