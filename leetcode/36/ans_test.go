package main

import "testing"

var tests = []struct {
	name  string
	board [][]byte
	want  bool
}{
	{
		name: "1",
		board: [][]byte{
			{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		want: true,
	},
	{
		name: "2",
		board: [][]byte{
			{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
			{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
			{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
			{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
			{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
			{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
			{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
			{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
			{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
		},
		want: false,
	},
	{
		name: "3",
		board: [][]byte{
			{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '3', '.', '.'},
			{'.', '.', '.', '1', '8', '.', '.', '.', '.'},
			{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '1', '.', '9', '7', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '3', '6', '.', '1', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
			{'.', '.', '.', '.', '.', '.', '.', '2', '.'},
		},
		want: true,
	},
}

func Test_isValidSudoku(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidSudokuBool(tt.board); got != tt.want {
				t.Errorf("isValidSudokuBool() = %v, want %v", got, tt.want)
			}

			if got := isValidSudokuBit(tt.board); got != tt.want {
				t.Errorf("isValidSudokuBit() = %v, want %v", got, tt.want)
			}
		})
	}
}
