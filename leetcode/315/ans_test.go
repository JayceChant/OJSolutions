package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	name string
	nums []int
	want []int
}{
	{
		name: "0",
		nums: []int{},
		want: []int{},
	},
	{
		name: "1element",
		nums: []int{1},
		want: []int{0},
	},
	{
		name: "1",
		nums: []int{5, 2, 6, 1},
		want: []int{2, 1, 1, 0},
	},
	{
		name: "full",
		nums: []int{81, 87, 47, 59, 81, 18, 25, 40, 56, 0, 94, 11, 62, 89, 28, 74, 11, 45, 37, 6, 95, 66, 28, 58, 47, 47, 87, 88, 90, 15, 41, 8, 87, 31, 29, 56, 37, 31, 85, 26, 13, 90, 94, 63, 33, 47, 78, 24, 59, 53, 57, 21, 89, 99, 0, 5, 88, 38, 3, 55, 51, 10, 5, 56, 66, 28, 61, 2, 83, 46, 63, 76, 2, 18, 47, 94, 77, 63, 96, 20, 23, 53, 37, 33, 41, 59, 33, 43, 91, 2, 78, 36, 46, 7, 40, 3, 52, 43, 5, 98},
		want: []int{79, 82, 49, 63, 77, 18, 23, 38, 55, 0, 83, 13, 60, 76, 22, 64, 13, 39, 30, 9, 76, 58, 20, 50, 38, 38, 61, 62, 64, 13, 31, 10, 58, 20, 19, 40, 24, 19, 52, 17, 11, 52, 53, 41, 17, 28, 44, 15, 35, 30, 33, 13, 42, 46, 0, 5, 39, 17, 3, 26, 23, 7, 4, 23, 27, 9, 23, 0, 27, 16, 21, 22, 0, 4, 15, 22, 19, 18, 20, 4, 4, 14, 7, 4, 7, 11, 4, 6, 10, 0, 8, 3, 5, 2, 2, 0, 2, 1, 0, 0},
	},
}

func Test_countSmaller(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSmallerBF(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmallerBF() = %v, want %v", got, tt.want)
			}

			if got := countSmallerMerge(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmallerMerge() = %v, want %v", got, tt.want)
			}

			if got := countSmallerAVL(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("countSmallerAVL() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_countSmallerBF(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			countSmallerBF(tt.nums)
		}
	}
}

func Benchmark_countSmallerMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			countSmallerMerge(tt.nums)
		}
	}
}

func Benchmark_countSmallerAVL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			countSmallerAVL(tt.nums)
		}
	}
}
