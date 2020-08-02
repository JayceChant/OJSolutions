package main

// binary search
// 0ms
// 2.2MB
func mySqrtBS(x int) int {
	low := 0
	high := x
	for {
		r := (low + high) / 2
		sqr := r * r
		if sqr == x {
			return r
		}
		if sqr < x {
			low = r + 1
		} else {
			high = r - 1
		}
		if low > high {
			break
		}
	}
	return low - 1
}

// Newton iteration
// 4ms
// 2.2MB
func mySqrt(x int) int {
	if x == 0 {
		return 0
	}
	r := x
	for {
		d := (r - x/r) / 2
		if d == 0 {
			break
		}
		r -= d
	}
	if r*r > x {
		return r - 1
	}
	return r
}
