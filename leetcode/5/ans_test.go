package main

import "testing"

var tests = []struct {
	name string
	s    string
	want []string
}{
	{
		name: "00",
		s:    "",
		want: []string{
			"",
		},
	},
	{
		name: "0",
		s:    "abcd",
		want: []string{
			"a",
			"b",
			"c",
			"d",
		},
	},
	{
		name: "1",
		s:    "babad",
		want: []string{
			"bab",
			"aba",
		},
	},
	{
		name: "2",
		s:    "cbbd",
		want: []string{
			"bb",
		},
	},
	{
		name: "full_even",
		s:    "bbbb",
		want: []string{
			"bbbb",
		},
	},
	{
		name: "full_odd",
		s:    "baaab",
		want: []string{
			"baaab",
		},
	},
	{
		name: "edge1",
		s:    "abb",
		want: []string{
			"bb",
		},
	},
	{
		name: "big_case",
		s:    "ababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababcbabababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababa",
		want: []string{
			"ababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababcbabababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababababa",
		},
	},
}

func contains(want []string, s string) bool {
	for _, w := range want {
		if s == w {
			return true
		}
	}
	return false
}

func Test_longestPalindrome(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.s); !contains(tt.want, got) {
				t.Errorf("longestPalindrome() = %v, want %v", got, tt.want)
			}
			if got := longestPalindrome2(tt.s); !contains(tt.want, got) {
				t.Errorf("longestPalindrome2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_longestPalindrome(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			longestPalindrome(tt.s)
		}
	}
}

func Benchmark_longestPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			longestPalindrome2(tt.s)
		}
	}
}
