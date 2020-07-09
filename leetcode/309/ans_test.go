package main

import "testing"

var tests = []struct {
	name   string
	prices []int
	want   int
}{
	{
		name:   "0",
		prices: []int{},
		want:   0,
	},
	{
		name:   "1",
		prices: []int{1, 2, 3, 0, 2},
		want:   3,
	},
	{
		name:   "fail",
		prices: []int{2, 1, 4},
		want:   3,
	},
}

func Test_maxProfit(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProfitN2(tt.prices); got != tt.want {
				t.Errorf("maxProfitN2() = %v, want %v", got, tt.want)
			}

			if got := maxProfit(tt.prices); got != tt.want {
				t.Errorf("maxProfit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_maxProfitN2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			maxProfitN2(tt.prices)
		}
	}
}

func Benchmark_maxProfit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			maxProfit(tt.prices)
		}
	}
}
