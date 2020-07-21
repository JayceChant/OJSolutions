package main

import "testing"

var tests = []struct {
	name string
	nums []int
	want int
}{
	{
		name: "1",
		nums: []int{5, 3, 2, 4},
		want: 0,
	},
	{
		name: "2",
		nums: []int{1, 5, 0, 10, 14},
		want: 1,
	},
	{
		name: "3",
		nums: []int{6, 6, 0, 1, 1, 4, 6},
		want: 2,
	},
	{
		name: "4",
		nums: []int{1, 5, 6, 14, 15},
		want: 1,
	},
	{
		name: "fail",
		nums: []int{82, 81, 95, 75, 20},
		want: 1,
	},
}

func Test_minDifference(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDifference(tt.nums); got != tt.want {
				t.Errorf("minDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
