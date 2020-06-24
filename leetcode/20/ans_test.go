package main

import "testing"

var tests = []struct {
	name string
	s    string
	want bool
}{
	{
		name: "0",
		s:    "",
		want: true,
	},
	{
		name: "1",
		s:    "()",
		want: true,
	},
	{
		name: "2",
		s:    "()[]{}",
		want: true,
	},
	{
		name: "3",
		s:    "(]",
		want: false,
	},
	{
		name: "4",
		s:    "([)]",
		want: false,
	},
	{
		name: "5",
		s:    "{[]}",
		want: true,
	},
	{
		name: "6",
		s:    "((()[]{}",
		want: false,
	},
	{
		name: "7",
		s:    "((()[])){}",
		want: true,
	},
	{
		name: "8",
		s:    "((()[])))){}",
		want: false,
	},
}

func Test_isValid(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidMap(tt.s); got != tt.want {
				t.Errorf("isValidMap() = %v, want %v", got, tt.want)
			}

			if got := isValid(tt.s); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_isValidMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			isValidMap(tt.s)
		}
	}
}

func Benchmark_isValid(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			isValid(tt.s)
		}
	}
}
