package main

// 36ms
// 6.1MB
func rangeBitwiseAnd1(m int, n int) int {
	if m&n == 0 {
		return 0
	}

	mh := highbit(m)
	nh := highbit(n)
	if mh != nh {
		return 0
	}

	return mh | rangeBitwiseAnd1(m-mh, n-nh)
}

func highbit(a int) int {
	var b int
	for {
		b = a & (a - 1)
		if b == 0 {
			return a
		}
		a = b
	}
}

// 8ms
// 6.1MB
func rangeBitwiseAnd2(m int, n int) int {
	if m&n == 0 {
		return 0
	}
	for n > m {
		n &= (n - 1)
	}
	return n
}

// 8ms
// 6.1MB
func rangeBitwiseAnd(m int, n int) int {
	if m == n {
		return n
	}
	if m&n == 0 {
		return 0
	}
	low := 1
	high := 31
	xor := m ^ n
	for low < high {
		mid := (low + high) / 2
		if xor&(-1<<mid) == 0 {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return n & (-1 << high)
}
