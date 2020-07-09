package main

import "testing"

var tests = []struct {
	name   string
	matrix [][]int
	k      int
	want   int
}{
	{
		name: "1",
		matrix: [][]int{
			{1},
		},
		k:    1,
		want: 1,
	},
	{
		name: "2",
		matrix: [][]int{
			{1, 5, 9},
			{10, 11, 13},
			{12, 13, 15},
		},
		k:    8,
		want: 13,
	},
}

func Test_kthSmallest(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthSmallest(tt.matrix, tt.k); got != tt.want {
				t.Errorf("kthSmallest() = %v, want %v", got, tt.want)
			}
		})
	}
}
