package main

import (
	"reflect"
	"testing"
)

func Test_maxRectangle(t *testing.T) {
	tests := []struct {
		name  string
		words []string
		want  []string
	}{
		{
			name:  "1",
			words: []string{"this", "real", "hard", "trh", "hea", "iar", "sld"},
			want:  []string{"this", "real", "hard"},
		},
		{
			name:  "2",
			words: []string{"aa"},
			want:  []string{"aa", "aa"},
		},
		{
			name:  "3",
			words: []string{"this", "real", "hard", "trh", "hea", "iar", "sld", "aa"},
			want:  []string{"this", "real", "hard"},
		},
		{
			name:  "4",
			words: []string{"this", "reel", "hard", "trh", "hea", "iar", "sld", "aa"},
			want:  []string{"aa", "aa"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxRectangle(tt.words); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("maxRectangle() = %v, want %v", got, tt.want)
			}
		})
	}
}
