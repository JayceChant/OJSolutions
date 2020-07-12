package main

import "testing"

var tests = []struct {
	name string
	nums []int
	want int
}{
	{
		name: "1num",
		nums: []int{1},
		want: 0,
	},
	{
		name: "1",
		nums: []int{1, 2, 3, 1, 1, 3},
		want: 4,
	},
	{
		name: "2",
		nums: []int{1, 1, 1, 1},
		want: 6,
	},
	{
		name: "3",
		nums: []int{1, 2, 3},
		want: 0,
	},
}

func Test_numIdenticalPairs(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIdenticalPairs(tt.nums); got != tt.want {
				t.Errorf("numIdenticalPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
