package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name   string
	nums   []int
	target int
	want   []int
}{
	{
		name:   "1",
		nums:   []int{1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 4},
		target: 2,
		want:   []int{1, 7},
	},
	{
		name:   "2",
		nums:   []int{1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 4},
		target: 0,
		want:   []int{-1, -1},
	},
	{
		name:   "3",
		nums:   []int{1, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 4},
		target: 5,
		want:   []int{-1, -1},
	},
	{
		name:   "4",
		nums:   []int{2, 2},
		target: 3,
		want:   []int{-1, -1},
	},
}

func Test_searchRange(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchRange(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("searchRange() = %v, want %v", got, tt.want)
			}
		})
	}
}
