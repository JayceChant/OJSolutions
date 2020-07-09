package main

import "testing"

func Test_numSubmat(t *testing.T) {
	tests := []struct {
		name string
		mat  [][]int
		want int
	}{
		{
			name: "min0",
			mat: [][]int{
				{0},
			},
			want: 0,
		},
		{
			name: "min1",
			mat: [][]int{
				{1},
			},
			want: 1,
		},
		{
			name: "1",
			mat: [][]int{
				{1, 0, 1},
				{1, 1, 0},
				{1, 1, 0},
			},
			want: 13,
		},
		{
			name: "2",
			mat: [][]int{
				{0, 1, 1, 0},
				{0, 1, 1, 1},
				{1, 1, 1, 0},
			},
			want: 24,
		},
		{
			name: "3",
			mat: [][]int{
				{1, 1, 1, 1, 1, 1},
			},
			want: 21,
		},
		{
			name: "4",
			mat: [][]int{
				{1, 0, 1},
				{0, 1, 0},
				{1, 0, 1},
			},
			want: 5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numSubmat(tt.mat); got != tt.want {
				t.Errorf("numSubmat() = %v, want %v", got, tt.want)
			}
		})
	}
}
