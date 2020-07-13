package main

import "testing"

var tests = []struct {
	name string
	nums []int
	want [][]int
}{
	{
		name: "1",
		nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
		want: [][]int{
			{-2, 1, 1, 4, 4, 5, 6, 6, 6},
			{0, 1, 1, 4, 4, 5, 6, 6, 6},
			{0, 0, -3, 4, 4, 5, 6, 6, 6},
			{0, 0, 0, 4, 4, 5, 6, 6, 6},
			{0, 0, 0, 0, -1, 2, 3, 3, 4},
			{0, 0, 0, 0, 0, 2, 3, 3, 4},
			{0, 0, 0, 0, 0, 0, 1, 1, 4},
			{0, 0, 0, 0, 0, 0, 0, -5, 4},
			{0, 0, 0, 0, 0, 0, 0, 0, 4},
		},
	},
}

func Test_treeNode_maxSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			size := len(tt.nums)
			tree := buildTree(tt.nums, 0, size-1)
			for st := 0; st < size; st++ {
				for ed := st; ed < size; ed++ {
					if got := tree.maxSum(st, ed); got != tt.want[st][ed] {
						t.Errorf("maxSum expect=%d, got=%d", tt.want[st][ed], got)
					}
				}
			}
		})
	}
}

func Benchmark_maxSubArray(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			size := len(tt.nums)
			for st := 0; st < size; st++ {
				for ed := st + 1; ed <= size; ed++ {
					// [st, ed)
					maxSubArray(tt.nums[st:ed])
				}
			}
		}
	}
}

func Benchmark_maxSubArraySegmentTree(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			size := len(tt.nums)
			tree := buildTree(tt.nums, 0, size-1)
			for st := 0; st < size; st++ {
				for ed := st; ed < size; ed++ {
					tree.maxSum(st, ed)
				}
			}
		}
	}
}

var benchmarkData = []int{81, 87, 47, 59, 81, 18, 25, 40, 56, 0, 94, 11, 62, 89, 28, 74, 11, 45, 37, 6, 95, 66, 28, 58, 47, 47, 87, 88, 90, 15, 41, 8, 87, 31, 29, 56, 37, 31, 85, 26, 13, 90, 94, 63, 33, 47, 78, 24, 59, 53, 57, 21, 89, 99, 0, 5, 88, 38, 3, 55, 51, 10, 5, 56, 66, 28, 61, 2, 83, 46, 63, 76, 2, 18, 47, 94, 77, 63, 96, 20, 23, 53, 37, 33, 41, 59, 33, 43, 91, 2, 78, 36, 46, 7, 40, 3, 52, 43, 5, 98}

func Benchmark_maxSubArrayOneCase(b *testing.B) {
	size := len(benchmarkData)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for st := 0; st < size; st++ {
			for ed := st + 1; ed <= size; ed++ {
				// [st, ed)
				maxSubArray(benchmarkData[st:ed])
			}
		}
	}
}

func Benchmark_maxSubArraySegmentTreeOneCase(b *testing.B) {
	size := len(benchmarkData)
	tree := buildTree(benchmarkData, 0, size-1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for st := 0; st < size; st++ {
			for ed := st; ed < size; ed++ {
				tree.maxSum(st, ed)
			}
		}
	}
}
