package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name string
	nums []int
	want [][]int
}{
	{
		name: "0ans",
		nums: []int{-1, 0},
		want: [][]int{},
	},
	{
		name: "1",
		nums: []int{-1, 0, 1, 2, -1, -4},
		want: [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
		},
	},
	{
		name: "all0",
		nums: []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		want: [][]int{
			{0, 0, 0},
		},
	},
	{
		name: "dupAns",
		nums: []int{-1, -1, -1, -1, 0, 0, 0, 0, 1, 1, 1, 1, 2, 2, 2, 2, 2},
		want: [][]int{
			{-1, -1, 2},
			{-1, 0, 1},
			{0, 0, 0},
		},
	},
	{
		name: "notenough0",
		nums: []int{0, 0},
		want: [][]int{},
	},
}

func Test_threeSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := threeSumMap(tt.nums); !reflect.DeepEqual(got, tt.want) {
			// 	t.Errorf("threeSumMap() = %v, want %v", got, tt.want)
			// }
			if got := threeSum(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("threeSum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_threeSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			threeSum(tt.nums)
		}
	}
}

func Benchmark_threeSumMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			threeSumMap(tt.nums)
		}
	}
}

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
