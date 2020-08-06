package main

import (
	"reflect"
	"sort"
	"testing"
)

// const (
// 	str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
// )

// func Benchmark_reverse1(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		reverse1(str)
// 	}
// }

// func Benchmark_reverse2(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		reverse2(str)
// 	}
// }

func Test_palindromePairs(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  [][]int
	}{
		{
			name:  "fail with empty",
			words: []string{"a", ""},
			want: sortPairs([][]int{
				{0, 1},
				{1, 0},
			}),
		},
		{
			name:  "fail2 with empty",
			words: []string{"a", "abc", "aba", ""},
			want: sortPairs([][]int{
				{0, 3},
				{3, 0},
				{2, 3},
				{3, 2},
			}),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortPairs(palindromePairs(tt.words)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("palindromePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func sortPairs(pairs [][]int) [][]int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][0] < pairs[j][0] || (pairs[i][0] == pairs[j][0] && pairs[i][1] < pairs[j][1])
	})
	return pairs
}
