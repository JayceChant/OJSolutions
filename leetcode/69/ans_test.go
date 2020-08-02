package main

import (
	"fmt"
	"math"
	"testing"
)

const (
	testRange = 20
)

func Test_mySqrt(t *testing.T) {
	for x := 0; x <= testRange; x++ {
		t.Run(fmt.Sprint(x), func(t *testing.T) {
			want := int(math.Sqrt(float64(x)))
			if got := mySqrt(x); got != want {
				t.Errorf("mySqrt() = %v, want %v", got, want)
			}
		})
	}
}
