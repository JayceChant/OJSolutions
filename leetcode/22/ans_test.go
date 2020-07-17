package main

import (
	"fmt"
	"testing"
)

var tests = []struct {
	name string
	n    int
	want int
}{
	{
		name: "1",
		n:    3,
		want: 5,
	},
}

func Test_generateParenthesis(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateParenthesis(tt.n); !areVallid(got, tt.want) {
				t.Errorf("generateParenthesis() = %v, want %v answers", got, tt.want)
			}

			if got := generateParenthesisMemoDAC(tt.n); !areVallid(got, tt.want) {
				t.Errorf("generateParenthesisMemoDAC() = %v, want %v answers", got, tt.want)
			}
		})
	}
}

const (
	benchmarkN = 14
)

func Benchmark_generateParenthesis(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generateParenthesis(benchmarkN)
	}
}

func Benchmark_generateParenthesisMemoDAC(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := 0; n <= benchmarkN; n++ {
			generateParenthesisMemoDAC(n)
		}
	}
}

func areVallid(strs []string, n int) bool {
	if len(strs) != n {
		fmt.Printf("has %v answers, want %v\n", len(strs), n)
		return false
	}

	strSet := make(map[string]bool)
	for _, str := range strs {
		if !isValid(str) {
			fmt.Println(str, "is not Valid")
			return false
		}
		strSet[str] = true
	}

	if len(strSet) != n {
		fmt.Printf("has %v unique answers, want %v\n", len(strs), n)
		return false
	}
	return true
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	half := len(s) / 2
	left := 0

	for _, char := range s {
		if char == ')' {
			if left == 0 {
				return false
			}
			left--
		} else {
			if left == half {
				return false
			}
			left++
		}
	}
	return left == 0
}
