package main

import "testing"

var tests2 = []struct {
	name  string
	n     int
	left  []int
	right []int
	want  int
}{
	{
		name:  "1",
		n:     4,
		left:  []int{4, 3},
		right: []int{0, 1},
		want:  4,
	},
	{
		name:  "2",
		n:     7,
		left:  []int{},
		right: []int{0, 1, 2, 3, 4, 5, 6, 7},
		want:  7,
	},
	{
		name:  "3",
		n:     7,
		left:  []int{0, 1, 2, 3, 4, 5, 6, 7},
		right: []int{},
		want:  7,
	},
	{
		name:  "4",
		n:     9,
		left:  []int{5},
		right: []int{4},
		want:  5,
	},
	{
		name:  "5",
		n:     6,
		left:  []int{6},
		right: []int{0},
		want:  6,
	},
}

func Test_getLastMoment(t *testing.T) {
	for _, tt := range tests2 {
		t.Run(tt.name, func(t *testing.T) {
			if got := getLastMoment(tt.n, tt.left, tt.right); got != tt.want {
				t.Errorf("getLastMoment() = %v, want %v", got, tt.want)
			}
		})
	}
}
