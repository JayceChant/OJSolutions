package main

import (
	"testing"
)

var tests = []struct {
	name string
	s    string
	p    string
	want bool
}{
	{
		name: "0v0",
		s:    "",
		p:    "",
		want: true,
	},
	{
		name: "0va",
		s:    "",
		p:    "a",
		want: false,
	},
	{
		name: "0va*",
		s:    "",
		p:    "a*",
		want: true,
	},
	{
		name: "av0",
		s:    "a",
		p:    "",
		want: false,
	},
	{
		name: "1",
		s:    "aa",
		p:    "a",
		want: false,
	},
	{
		name: "2",
		s:    "aa",
		p:    "a*",
		want: true,
	},
	{
		name: "3",
		s:    "ab",
		p:    ".*",
		want: true,
	},
	{
		name: "4",
		s:    "aab",
		p:    "c*a*b",
		want: true,
	},
	{
		name: "5",
		s:    "mississippi",
		p:    "mis*is*p*.",
		want: false,
	},
	// 篇幅关系，只提供题目的例子
	// 实际开发一定要多考虑不同的测试用例
	{
		name: "6",
		s:    "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab",
		p:    "caa*a*a*a*b*",
		want: false,
	},
	{
		name: "7",
		s:    "caaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa",
		p:    "c*aa*a*a*a*b",
		want: false,
	},
}

func TestIsMatch(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isMatch(tt.s, tt.p); got != tt.want {
				t.Errorf("isMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isMatchBT(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			isMatchBT(tt.s, tt.p)
		}
	}
}

func Benchmark_isMatchMemoRL(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			isMatchMemoRL(tt.s, tt.p)
		}
	}
}

func Benchmark_isMatchMemo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			isMatchMemo(tt.s, tt.p)
		}
	}
}

func Benchmark_isMatchDP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			isMatchDP(tt.s, tt.p)
		}
	}
}

func Benchmark_isMatchDP2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			isMatchDP2(tt.s, tt.p)
		}
	}
}
