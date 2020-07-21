package main

import (
	"reflect"
	"sort"
	"testing"
)

func Test_maxNumOfSubstrings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{
			name: "1",
			s:    "adefaddaccc",
			want: []string{"ccc", "e", "f"},
		},
		{
			name: "2",
			s:    "abbaccd",
			want: []string{"bb", "cc", "d"},
		},
		{
			name: "fail",
			s:    "abab",
			want: []string{"abab"},
		},
		{
			name: "fail2",
			s:    "abcdefghahbgcfed",
			want: []string{"abcdefghahbgcfed"},
		},
		{
			name: "fail3",
			s:    "ababa",
			want: []string{"ababa"},
		},
		{
			name: "fail4",
			s:    "bacabbbc",
			want: []string{"bacabbbc"},
		},
		{
			name: "fail5",
			s:    "abcbaaac",
			want: []string{"abcbaaac"},
		},
		{
			name: "edge",
			s:    "abcdcdhahbcz",
			want: []string{"abcdcdhahbc", "z"},
		},
		{
			name: "edge2",
			s:    "abcdcdghahbcz",
			want: []string{"g", "z"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxNumOfSubstrings(tt.s)
			sort.Strings(got)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxNumOfSubstrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
