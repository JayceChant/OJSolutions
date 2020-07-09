package main

import "testing"

var tests = []struct {
	name    string
	s       string
	numRows int
	want    string
}{
	{
		name:    "0str",
		s:       "",
		numRows: 1,
		want:    "",
	},
	{
		name:    "1row",
		s:       "abcdefg",
		numRows: 1,
		want:    "abcdefg",
	},
	{
		name:    "1",
		s:       "LEETCODEISHIRING",
		numRows: 3,
		want:    "LCIRETOESIIGEDHN",
	},
	{
		name:    "2",
		s:       "LEETCODEISHIRING",
		numRows: 4,
		want:    "LDREOEIIECIHNTSG",
	},
	{
		name:    "3",
		s:       "PAYPALISHIRING",
		numRows: 3,
		want:    "PAHNAPLSIIGYIR",
	},
	{
		name:    "4",
		s:       "LEETCODEISHIRINGPAYPALISHIRINGLCIRETOESIIGEDHNPAHNAPLSIIGYIRLDREOEIIECIHNTSG",
		numRows: 7,
		want:    "LRHOHLNEIISITEANRDHTEHNIRESPAIRISTSGLIRINPYECGCIPANIIHLGOEOEAPGCGDSIEIDYLEII",
	},
}

func Test_convert(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convert(tt.s, tt.numRows); got != tt.want {
				t.Errorf("convert() = %v, want %v", got, tt.want)
			}

			if got := convert2(tt.s, tt.numRows); got != tt.want {
				t.Errorf("convert2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_convert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			convert(tt.s, tt.numRows)
		}
	}
}

func Benchmark_convert2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			convert2(tt.s, tt.numRows)
		}
	}
}
