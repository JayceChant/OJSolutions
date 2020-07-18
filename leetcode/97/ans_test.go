package main

import "testing"

var tests = []struct {
	name string
	s1   string
	s2   string
	s3   string
	want bool
}{
	{
		name: "bothempty",
		s1:   "",
		s2:   "",
		s3:   "",
		want: true,
	},
	{
		name: "s1empty-true",
		s1:   "",
		s2:   "abc",
		s3:   "abc",
		want: true,
	},
	{
		name: "s2empty-true",
		s1:   "abc",
		s2:   "",
		s3:   "abc",
		want: true,
	},
	{
		name: "s1empty-false",
		s1:   "",
		s2:   "abc",
		s3:   "aba",
		want: false,
	},
	{
		name: "s2empty-false",
		s1:   "abc",
		s2:   "",
		s3:   "aba",
		want: false,
	},
	{
		name: "1",
		s1:   "aabcc",
		s2:   "dbbca",
		s3:   "aadbbcbcac",
		want: true,
	},
	{
		name: "2",
		s1:   "aabcc",
		s2:   "dbbca",
		s3:   "aadbbbaccc",
		want: false,
	},
	{
		name: "edge1",
		s1:   "a",
		s2:   "b",
		s3:   "ab",
		want: true,
	},
	{
		name: "edge2",
		s1:   "aa",
		s2:   "ab",
		s3:   "aaba",
		want: true,
	},
	{
		name: "big",
		s1:   "bbbbbabbbbabaababaaaabbababbaaabbabbaaabaaaaababbbababbbbbabbbbababbabaabababbbaabababababbbaaababaa",
		s2:   "babaaaabbababbbabbbbaabaabbaabbbbaabaaabaababaaaabaaabbaaabaaaabaabaabbbbbbbbbbbabaaabbababbabbabaab",
		s3:   "babbbabbbaaabbababbbbababaabbabaabaaabbbbabbbaaabbbaaaaabbbbaabbaaabababbaaaaaabababbababaababbababbbababbbbaaaabaabbabbaaaaabbabbaaaabbbaabaaabaababaababbaaabbbbbabbbbaabbabaabbbbabaaabbababbabbabbab",
		want: false,
	},
}

func Test_isInterleave(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isInterleave(tt.s1, tt.s2, tt.s3); got != tt.want {
				t.Errorf("isInterleave() = %v, want %v", got, tt.want)
			}
		})
	}
}
