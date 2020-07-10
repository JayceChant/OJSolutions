package main

import "testing"

var tests = []struct {
	name string
	arr  []int
	want bool
}{
	{
		name: "edge",
		arr:  []int{1, 1},
		want: true,
	},
	{
		name: "1",
		arr:  []int{3, 5, 1},
		want: true,
	},
	{
		name: "1",
		arr:  []int{1, 2, 4},
		want: false,
	},
}

func Test_canMakeArithmeticProgression(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canMakeArithmeticProgression(tt.arr); got != tt.want {
				t.Errorf("canMakeArithmeticProgression() = %v, want %v", got, tt.want)
			}
		})
	}
}
