package main

import "testing"

var tests = []struct {
	name    string
	dungeon [][]int
	want    int
}{
	{
		name: "0",
		dungeon: [][]int{
			{100},
		},
		want: 1,
	},
	{
		name: "1",
		dungeon: [][]int{
			{-2, -3, 3},
			{-5, -10, 1},
			{10, 30, -5},
		},
		want: 7,
	},
	{
		name: "edge",
		dungeon: [][]int{
			{-2, -3, 3},
			{-5, -10, 1},
			{-50, 200, -100},
		},
		want: 16,
	},
}

func Test_calculateMinimumHP(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateMinimumHPBS(tt.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHPBS() = %v, want %v", got, tt.want)
			}

			if got := calculateMinimumHPBS2(tt.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHPBS2() = %v, want %v", got, tt.want)
			}

			if got := calculateMinimumHPDP(tt.dungeon); got != tt.want {
				t.Errorf("calculateMinimumHPDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_calculateMinimumHPBS(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			// clone test data
			myMat := make([][]int, len(tt.dungeon))
			for j := 0; j < len(myMat); j++ {
				myMat[j] = make([]int, len(tt.dungeon[j]))
				copy(myMat[j], tt.dungeon[j])
			}

			calculateMinimumHPBS(myMat)
		}
	}
}

func Benchmark_calculateMinimumHPBS2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			// clone test data
			myMat := make([][]int, len(tt.dungeon))
			for j := 0; j < len(myMat); j++ {
				myMat[j] = make([]int, len(tt.dungeon[j]))
				copy(myMat[j], tt.dungeon[j])
			}

			calculateMinimumHPBS2(myMat)
		}
	}
}

func Benchmark_calculateMinimumHPDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			// clone test data
			myMat := make([][]int, len(tt.dungeon))
			for j := 0; j < len(myMat); j++ {
				myMat[j] = make([]int, len(tt.dungeon[j]))
				copy(myMat[j], tt.dungeon[j])
			}

			calculateMinimumHPDP(myMat)
		}
	}
}
