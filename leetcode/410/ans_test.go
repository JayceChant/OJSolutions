package main

import "testing"

func Test_splitArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		m    int
		want int
	}{
		{
			name: "1",
			nums: []int{7, 2, 5, 10, 8},
			m:    2,
			want: 18,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := splitArray(tt.nums, tt.m); got != tt.want {
				t.Errorf("splitArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
