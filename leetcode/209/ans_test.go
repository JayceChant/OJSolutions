package main

import "testing"

var tests = []struct {
	name string
	s    int
	nums []int
	want int
}{
	{
		name: "0",
		s:    7,
		nums: []int{},
		want: 0,
	},
	{
		name: "1",
		s:    7,
		nums: []int{2, 3, 1, 2, 4, 3},
		want: 2,
	},
	{
		name: "big-s",
		s:    10000,
		nums: []int{2, 3, 1, 2, 4, 3},
		want: 0,
	},
	{
		name: "fail",
		s:    11,
		nums: []int{1, 2, 3, 4, 5},
		want: 3,
	},
}

func Test_minSubArrayLen(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.s, tt.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
