package main

// 0ms
// 2.1MB
func grayCodeMirror(n int) []int {
	size := 1 << n
	ans := make([]int, size)
	half := size >> 1
	for width := 1; width <= half; width <<= 1 {
		ed := width<<1 - 1
		for i := 0; i < width; i++ {
			ans[ed-i] = width + ans[i]
		}
	}
	return ans
}

// 0ms
// 2.1MB
func grayCode(n int) []int {
	size := 1 << n
	ans := make([]int, size)
	for i := 1; i < size; i++ {
		ans[i] = i ^ (i >> 1)
	}
	return ans
}
