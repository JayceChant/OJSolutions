package main

import "testing"

var tests = []struct {
	name string
	s    string
	want int
}{
	{
		name: "0",
		s:    "",
		want: 0,
	},
	{
		name: "space",
		s:    " ",
		want: 1,
	},
	{
		name: "1",
		s:    "abcabcbb",
		want: 3,
	},
	{
		name: "2",
		s:    "bbbbb",
		want: 1,
	},
	{
		name: "3",
		s:    "pwwkew",
		want: 3,
	},
	{
		name: "fail",
		s:    "abba",
		want: 2,
	},
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLongestSubstringMap(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstringMap() = %v, want %v", got, tt.want)
			}

			if got := lengthOfLongestSubstring(tt.s); got != tt.want {
				t.Errorf("lengthOfLongestSubstring() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_lengthOfLongestSubstringMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			lengthOfLongestSubstringMap(tt.s)
		}
	}
}

func Benchmark_lengthOfLongestSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			lengthOfLongestSubstring(tt.s)
		}
	}
}
