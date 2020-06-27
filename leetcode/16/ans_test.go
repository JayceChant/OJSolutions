package main

import (
	"reflect"
	"testing"
)

var sortTests = []struct {
	name string
	nums []int
	want []int
}{
	{
		name: "1",
		nums: []int{-1, 0, 1, 2, -1, -4},
		want: []int{-4, -1, -1, 0, 1, 2},
	},
}

func Test_quickSort(t *testing.T) {
	for _, tt := range sortTests {
		t.Run(tt.name, func(t *testing.T) {
			if quickSort(tt.nums); !reflect.DeepEqual(tt.nums, tt.want) {
				t.Errorf("quickSort() = %v, want %v", tt.nums, tt.want)
			}
		})
	}
}

var tests = []struct {
	name   string
	nums   []int
	target int
	want   int
}{
	{
		name:   "1",
		nums:   []int{-1, 2, 1, -4},
		target: 1,
		want:   2,
	},
	{
		name:   "fail",
		nums:   []int{1, 1, 1, 1},
		target: 0,
		want:   3,
	},
}

func Test_threeSumClosest(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := threeSumClosest(tt.nums, tt.target); got != tt.want {
				t.Errorf("threeSumClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
