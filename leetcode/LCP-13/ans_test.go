package main

import "testing"

func Test_minimalSteps(t *testing.T) {
	tests := []struct {
		name string
		maze []string
		want int
	}{
		{
			name: "no-M",
			maze: []string{
				"S#O",
				"...",
				"..T",
			},
			want: 4,
		},
		{
			name: "blocked-M",
			maze: []string{
				"S#M",
				"O##",
				"..T",
			},
			want: -1,
		},
		{
			name: "blocked-O",
			maze: []string{
				"S#O",
				"M##",
				"..T",
			},
			want: -1,
		},
		{
			name: "1",
			maze: []string{
				"S#O",
				"M..",
				"M.T",
			},
			want: 16,
		},
		{
			name: "2",
			maze: []string{
				"S.O",
				"...",
				"...",
				"O..",
				"M..",
				"T..",
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minimalSteps(tt.maze); got != tt.want {
				t.Errorf("minimalSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
