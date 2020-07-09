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
		name: "0v*",
		s:    "",
		p:    "*",
		want: true,
	},
	{
		name: "0va*",
		s:    "",
		p:    "a*",
		want: false,
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
		p:    "*",
		want: true,
	},
	{
		name: "3",
		s:    "cb",
		p:    "?a",
		want: false,
	},
	{
		name: "4",
		s:    "adceb",
		p:    "*a*b",
		want: true,
	},
	{
		name: "5",
		s:    "acdcb",
		p:    "a*c?b",
		want: false,
	},
	// 篇幅关系，只提供题目的例子
	// 实际开发一定要多考虑不同的测试用例
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
