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
		nums:   []int{2, 7, 11, 15},
		target: 9,
		want:   []int{0, 1},
	},
}

func Test_twoSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSum(tt.nums, tt.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
