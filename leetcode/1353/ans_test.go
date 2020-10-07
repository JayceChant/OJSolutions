package main

import "testing"

func Test_maxEvents(t *testing.T) {
	tests := []struct {
		name   string
		events [][]int
		want   int
	}{
		{
			name:   "1",
			events: [][]int{{1, 5}, {1, 5}, {1, 5}, {2, 3}, {2, 3}},
			want:   5,
		},
		{
			name:   "fail",
			events: [][]int{{27, 27}, {8, 10}, {9, 11}, {20, 21}, {25, 29}, {17, 20}, {12, 12}, {12, 12}, {10, 14}, {7, 7}, {6, 10}, {7, 7}, {4, 8}, {30, 31}, {23, 25}, {4, 6}, {17, 17}, {13, 14}, {6, 9}, {13, 14}},
			want:   18,
		},
		{
			name:   "fail2",
			events: [][]int{{27, 29}, {28, 32}, {3, 3}, {24, 25}, {7, 7}, {22, 25}, {14, 15}, {13, 17}, {1, 2}, {7, 7}, {10, 12}, {9, 13}, {21, 25}, {20, 21}, {20, 22}, {19, 20}, {27, 28}, {9, 9}, {21, 24}, {18, 21}, {6, 10}, {29, 30}, {22, 24}},
			want:   21,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEvents(tt.events); got != tt.want {
				t.Errorf("maxEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
