package main

import "testing"

var tests = []struct {
	name string
	n    int
	want bool
}{
	{
		name: "1",
		n:    1,
		want: true,
	},
	{
		name: "2",
		n:    2,
		want: false,
	},
	{
		name: "4",
		n:    4,
		want: true,
	},
	{
		name: "7",
		n:    7,
		want: false,
	},
	{
		name: "17",
		n:    17,
		want: false,
	},
	{
		name: "100",
		n:    100,
		want: true,
	},
	{
		name: "1000",
		n:    1000,
		want: true,
	},
}

func Test_winnerSquareGame(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// if got := winnerSquareGameBF(tt.n); got != tt.want {
			// 	t.Errorf("winnerSquareGameBF() = %v, want %v", got, tt.want)
			// }

			if got := winnerSquareGame(tt.n); got != tt.want {
				t.Errorf("winnerSquareGame() = %v, want %v", got, tt.want)
			}
		})
	}
}
