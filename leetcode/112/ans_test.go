package main

import "testing"

var tests = []struct {
	name string
	root *TreeNode
	sum  int
	want bool
}{
	{
		name: "fail",
		root: nil,
		sum:  0,
		want: false,
	},
}

func Test_hasPathSum(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasPathSum(tt.root, tt.sum); got != tt.want {
				t.Errorf("hasPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
