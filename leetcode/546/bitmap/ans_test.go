package main

import "testing"

func Test_removeBoxes(t *testing.T) {
	tests := []struct {
		name  string
		boxes []int
		want  int
	}{
		// {
		// 	name:  "1",
		// 	boxes: []int{1, 3, 2, 2, 2, 3, 4, 3, 1},
		// 	want:  23,
		// },
		// {
		// 	name:  "min-fail",
		// 	boxes: []int{4, 4, 6, 5, 8, 4, 8, 6, 9, 6, 2, 8, 6, 4},
		// 	want:  34,
		// },
		// {
		// 	name:  "fail",
		// 	boxes: []int{3, 8, 8, 5, 5, 3, 9, 2, 4, 4, 6, 5, 8, 4, 8, 6, 9, 6, 2, 8, 6, 4, 1, 9, 5, 3, 10, 5, 3, 3, 9, 8, 8, 6, 5, 3, 7, 4, 9, 6, 3, 9, 4, 3, 5, 10, 7, 6, 10, 7},
		// 	want:  136,
		// },
		{
			name:  "fail2",
			boxes: []int{1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2, 1, 2},
			want:  2550,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := removeBoxes(tt.boxes); got != tt.want {
				t.Errorf("removeBoxes() = %v, want %v", got, tt.want)
			}
		})
	}
}
