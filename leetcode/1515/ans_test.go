package main

import (
	"math"
	"testing"
)

var tests = []struct {
	name      string
	positions [][]int
	want      float64
}{
	{
		name:      "0",
		positions: [][]int{{0, 1}},
		want:      0.0,
	},
	{
		name:      "1",
		positions: [][]int{{0, 1}, {1, 0}, {1, 2}, {2, 1}},
		want:      4.0,
	},
	{
		name:      "2",
		positions: [][]int{{1, 1}, {3, 3}},
		want:      2.82843,
	}, {
		name:      "3",
		positions: [][]int{{1, 1}},
		want:      0.0,
	}, {
		name:      "4",
		positions: [][]int{{1, 1}, {0, 0}, {2, 0}},
		want:      2.73205,
	}, {
		name:      "5",
		positions: [][]int{{0, 1}, {3, 2}, {4, 5}, {7, 6}, {8, 9}, {11, 1}, {2, 12}},
		want:      32.94036,
	},
}

func Test_getMinDistSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getMinDistSum(tt.positions); math.Abs(tt.want-got) >= precision {
				t.Errorf("getMinDistSum() = %v, want %v", got, tt.want)
			}

			if got := getMinDistSum2(tt.positions); math.Abs(tt.want-got) >= precision {
				t.Errorf("getMinDistSum2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_getMinDistSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			getMinDistSum(tt.positions)
		}
	}
}

func Benchmark_getMinDistSum2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			getMinDistSum2(tt.positions)
		}
	}
}
