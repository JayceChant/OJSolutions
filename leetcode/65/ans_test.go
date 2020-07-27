package main

import "testing"

func Test_isNumber(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{
			s:    "0",
			want: true,
		},
		{
			s:    "0.1",
			want: true,
		},
		{
			s:    "abc",
			want: false,
		},
		{
			s:    "1 a",
			want: false,
		},
		{
			s:    "2e10",
			want: true,
		},
		{
			s:    "-90e3",
			want: true,
		},
		{
			s:    "1e",
			want: false,
		},
		{
			s:    "e3",
			want: false,
		},
		{
			s:    "6e-1",
			want: true,
		},
		{
			s:    "99e2.5",
			want: false,
		},
		{
			s:    "53.5e93",
			want: true,
		},
		{
			s:    " --6 ",
			want: false,
		},
		{
			s:    "-+3",
			want: false,
		},
		{
			s:    "95e54e53",
			want: false,
		},
		{
			s:    ".1",
			want: true,
		},
		{
			s:    "1.",
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := isNumber(tt.s); got != tt.want {
				t.Errorf("isNumber() = %v, want %v", got, tt.want)
			}

			// if got := isNumHelper(tt.s); got != tt.want {
			// 	t.Errorf("isNumHelper() = %v, want %v", got, tt.want)
			// }
		})
	}
}
