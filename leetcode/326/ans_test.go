package main

import "testing"

func Test_isPowerOfThree(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{
			name: "-3",
			n:    -3,
			want: false,
		},
		{
			name: "0",
			n:    0,
			want: false,
		},
		{
			name: "1",
			n:    1,
			want: true,
		},
		{
			name: "2",
			n:    2,
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfThree(tt.n); got != tt.want {
				t.Errorf("isPowerOfThree() = %v, want %v", got, tt.want)
			}
		})
	}
}
