package main

import "testing"

var tests = []struct {
	name string
	x    int
	want bool
}{
	{
		name: "<0",
		x:    -1,
		want: false,
	},
	{
		name: "0",
		x:    0,
		want: true,
	},
	{
		name: "1digit",
		x:    9,
		want: true,
	},
	{
		name: "10",
		x:    10,
		want: false,
	},
	{
		name: "odd",
		x:    121,
		want: true,
	},
	{
		name: "0 in odd",
		x:    1200021,
		want: true,
	},
	{
		name: "0 in false",
		x:    1000021,
		want: false,
	},
}

func Test_isPalindrome(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindromeItoa(tt.x); got != tt.want {
				t.Errorf("isPalindromeItoa() = %v, want %v", got, tt.want)
			}
			if got := isPalindrome(tt.x); got != tt.want {
				t.Errorf("isPalindrome() = %v, want %v", got, tt.want)
			}
			if got := isPalindromeMath(tt.x); got != tt.want {
				t.Errorf("isPalindromeMath() = %v, want %v", got, tt.want)
			}
		})
	}
}
