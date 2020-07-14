package main

import "testing"

var tests = []struct {
	name  string
	nums  []int
	notin int
}{
	{
		name:  "7",
		nums:  []int{4, 5, 6, 7, 0, 1, 2},
		notin: 8,
	},
	{
		name:  "8",
		nums:  []int{4, 5, 6, 7, 8, 0, 1, 2},
		notin: 9,
	},
	{
		name:  "9",
		nums:  []int{4, 5, 6, 7, 8, 0, 1, 2, 3},
		notin: 10,
	},
	{
		name:  "10",
		nums:  []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		notin: 11,
	},
}

func Test_search(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.notin); got != -1 {
				t.Errorf("search() = %v, want -1", got)
			}
			for i := 0; i < len(tt.nums); i++ {
				if got := search(tt.nums, tt.nums[i]); got != i {
					t.Errorf("search() = %v, want %v", got, i)
				}
			}
		})
	}
}
