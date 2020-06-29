package main

import "testing"

var tests = []struct {
	name string
	nums []int
	k    int
	want int
}{
	{
		name: "1",
		nums: []int{3, 2, 1, 5, 6, 4},
		k:    2,
		want: 5,
	},
	{
		name: "2",
		nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
		k:    4,
		want: 4,
	},
	{
		name: "fail",
		nums: []int{99, 99},
		k:    1,
		want: 99,
	},
}

func Test_findKthLargest(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := findKthLargest(tt.nums, tt.k); got != tt.want {
			// 	t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			// }

			if got := findKthLargestHeap(tt.nums, tt.k); got != tt.want {
				t.Errorf("findKthLargestHeap() = %v, want %v", got, tt.want)
			}
		})
	}
}
