package main

import "testing"

var tests = []struct {
	name  string
	nums1 []int
	nums2 []int
	want  float64
}{
	// 题目有写：You may assume nums1 and nums2 cannot be both empty.
	// 但是生产环境还是的把这个测试用例加上
	// {
	// 	name:  "0v0",
	// 	nums1: []int{},
	// 	nums2: []int{},
	// 	want:  0.0,
	// },
	{
		name:  "0v1",
		nums1: []int{},
		nums2: []int{5},
		want:  5.0,
	},
	{
		name:  "2v0",
		nums1: []int{1, 8},
		nums2: []int{},
		want:  4.5,
	},
	{
		name:  "1",
		nums1: []int{1, 3},
		nums2: []int{2},
		want:  2.0,
	},
	{
		name:  "2",
		nums1: []int{1, 2},
		nums2: []int{3, 4},
		want:  2.5,
	},
	{
		name:  "3",
		nums1: []int{1, 2, 3, 4, 5},
		nums2: []int{6, 7, 8, 9},
		want:  5.0,
	},
	{
		name:  "4",
		nums1: []int{1, 3, 5, 7, 9},
		nums2: []int{2, 4, 6, 8, 10},
		want:  5.5,
	},
	{
		name:  "5",
		nums1: []int{1, 3, 5, 7},
		nums2: []int{4, 6, 8, 10},
		want:  5.5,
	},
	{
		name:  "6",
		nums1: []int{1, 3, 9, 9},
		nums2: []int{2, 4, 6, 9, 9, 9, 9},
		want:  9.0,
	},
}

func TestFindMedianSortedArrays(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMedianSortedArrays(tt.nums1, tt.nums2); got != tt.want {
				t.Errorf("findMedianSortedArrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
