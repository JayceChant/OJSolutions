package main

// 0ms, 100%
// 2MB, 100%
func plusOne(digits []int) []int {
	i := len(digits) - 1
	carry := true
	for i >= 0 && carry {
		digits[i]++
		if digits[i] >= 10 {
			digits[i] %= 10
		} else {
			carry = false
		}
		i--
	}
	if carry {
		return append([]int{1}, digits...)
	}
	return digits
}
