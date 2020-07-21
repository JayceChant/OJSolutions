package main

import (
	"reflect"
	"testing"
)

func Test_dedup(t *testing.T) {
	tests := []struct {
		name   string
		sorted []int
		want   []int
	}{
		{
			name:   "nodup",
			sorted: []int{1, 2, 3, 4, 5, 6, 7, 8},
			want:   []int{1, 2, 3, 4, 5, 6, 7, 8},
		},
		{
			name:   "alldup",
			sorted: []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1},
			want:   []int{1},
		},
		{
			name:   "1",
			sorted: []int{1, 2, 3, 3, 3, 3, 4, 5, 5, 5, 6, 6, 6, 7, 7, 7, 8, 8, 8, 9, 9, 10, 11},
			want:   []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := dedup(tt.sorted); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dedup() = %v, want %v", got, tt.want)
			}
		})
	}
}
