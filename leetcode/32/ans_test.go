package main

import "testing"

var tests = []struct {
	name string
	s    string
	want int
}{
	{
		name: "1",
		s:    "(()",
		want: 2,
	},
	{
		name: "2",
		s:    ")()())",
		want: 4,
	},
	{
		name: "3",
		s:    "(()(())())(((()(())())",
		want: 10,
	},
	{
		name: "4",
		s:    "(((()(())())(()(())())",
		want: 20,
	},
	{
		name: "fail",
		s:    ")()())()()(",
		want: 4,
	},
	{
		name: "fail2",
		s:    "))))())()()(()",
		want: 4,
	},
}

func Test_longestValidParentheses(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestValidParentheses(tt.s); got != tt.want {
				t.Errorf("longestValidParentheses() = %v, want %v", got, tt.want)
			}
			if got := longestValidParenthesesDP(tt.s); got != tt.want {
				t.Errorf("longestValidParenthesesDP() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_longestValidParentheses(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			longestValidParentheses(tt.s)
		}
	}
}

func Benchmark_longestValidParenthesesDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			longestValidParenthesesDP(tt.s)
		}
	}
}
