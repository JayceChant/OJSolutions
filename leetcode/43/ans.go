package main

import (
	"fmt"
	"strings"
)

// 0ms, 100%
// 2.4MB, 53.33%
func multiplyInt(str1 string, str2 string) string {
	if str1 == "0" || str2 == "0" {
		return "0"
	}
	len1 := len(str1)
	len2 := len(str2)
	nums1 := parse(str1, len1)
	nums2 := parse(str2, len2)
	res := make([]int, (len1+len2+3)/4)
	for i1, num1 := range nums1 {
		for i2, num2 := range nums2 {
			i := i1 + i2
			res[i] += num1 * num2
			for res[i] >= 10000 {
				res[i+1] += res[i] / 10000
				res[i] %= 10000
				i++
			}
		}
	}

	sb := strings.Builder{}
	i := len(res) - 1
	if res[i] == 0 {
		i--
	}
	sb.WriteString(fmt.Sprint(res[i]))
	i--
	for ; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%04d", res[i]))
	}
	return sb.String()
}

func parse(str string, length int) []int {
	nums := make([]int, (length+3)/4)
	start := length - 4
	end := length
	for {
		if start <= -4 {
			break
		}

		if start < 0 {
			start = 0
		}
		num := 0
		for i := start; i < end; i++ {
			num = num*10 + int(str[i]-'0')
		}
		// store each 4 digits per num
		// nums from low to high
		nums[(length-1-start)/4] = num
		end = start
		start -= 4
	}
	return nums
}

// 0ms
// 2.3MB
func multiply32(str1 string, str2 string) string {
	if str1 == "0" || str2 == "0" {
		return "0"
	}
	len1 := len(str1)
	len2 := len(str2)
	nums1 := parse32(str1, len1)
	nums2 := parse32(str2, len2)
	res := make([]uint32, (len1+len2+8)/9)
	for i1, num1 := range nums1 {
		for i2, num2 := range nums2 {
			i := i1 + i2
			sum := uint64(res[i]) + uint64(num1)*uint64(num2)
			for sum >= 1000000000 {
				res[i] = uint32(sum % 1000000000)
				sum = uint64(res[i+1]) + sum/1000000000
				i++
			}
			res[i] = uint32(sum)
		}
	}

	sb := strings.Builder{}
	i := len(res) - 1
	if res[i] == 0 {
		i--
	}
	sb.WriteString(fmt.Sprint(res[i]))
	i--
	for ; i >= 0; i-- {
		sb.WriteString(fmt.Sprintf("%09d", res[i]))
	}
	return sb.String()
}

func parse32(str string, length int) []uint32 {
	nums := make([]uint32, (length+8)/9)
	start := length - 9
	end := length
	for {
		if start <= -9 {
			break
		}

		if start < 0 {
			start = 0
		}
		var num uint32
		for i := start; i < end; i++ {
			num = num*10 + uint32(str[i]-'0')
		}
		// store each 9 digits per num
		// nums from low to high
		nums[(length-1-start)/9] = num
		end = start
		start -= 9
	}
	return nums
}

// It seems I misunderstood the question.
// Not only can't you use libraries (for big-int or str-conv),
// you can't even convert to integers on your own.
// 0ms, 100%
// 2.3MB, 76.92%
func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	len1 := len(num1)
	len2 := len(num2)
	lenRes := len1 + len2
	res := make([]byte, lenRes)

	for i1 := len1 - 1; i1 >= 0; i1-- {
		n1 := num1[i1] - '0'
		for i2 := len2 - 1; i2 >= 0; i2-- {
			i := i1 + i2 + 1
			d := res[i] + n1*(num2[i2]-'0')
			for d >= 10 {
				res[i] = d % 10
				i--
				d = res[i] + d/10
			}
			res[i] = d
		}
	}

	for i := lenRes - 1; i >= 0; i-- {
		res[i] += '0'
	}
	i := 0
	if res[i] == '0' {
		i++
	}
	return string(res[i:])
}
