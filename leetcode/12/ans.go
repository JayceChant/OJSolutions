package main

import "strings"

// 12ms, 12.46% (min 0ms)
// 3.4MB, 100%
func intToRoman1(num int) string {
	sb := new(strings.Builder)

	for i := 0; i < num/1000; i++ {
		sb.WriteByte('M')
	}
	num %= 1000
	intToRomanImpl(sb, num/100, 'M', 'D', 'C')
	num %= 100
	intToRomanImpl(sb, num/10, 'C', 'L', 'X')
	num %= 10
	intToRomanImpl(sb, num, 'X', 'V', 'I')
	return sb.String()
}

func intToRomanImpl(sb *strings.Builder, num int, ten byte, five byte, one byte) {
	if num == 9 {
		sb.WriteByte(one)
		sb.WriteByte(ten)
		return
	}
	if num == 4 {
		sb.WriteByte(one)
		sb.WriteByte(five)
		return
	}
	if num >= 5 {
		sb.WriteByte(five)
		num -= 5
	}
	for i := 0; i < num; i++ {
		sb.WriteByte(one)
	}
}

var (
	nums   = []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	romans = []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
)

// 4ms 96.32%
func intToRoman(num int) string {
	sb := strings.Builder{}
	i := 0
	for num > 0 {
		for num < nums[i] {
			i++
		}
		for num >= nums[i] {
			sb.WriteString(romans[i])
			num -= nums[i]
		}
	}
	return sb.String()
}
