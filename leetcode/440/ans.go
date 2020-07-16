package main

import (
	"fmt"
	"strconv"
)

var (
	// bases[i] means the base of numbers with i tailing 0
	// e.g. 100 has 2 tailing 0, bases[2] = 100
	bases = []int{
		1,
		10,
		100,
		1000,
		10000,
		100000,
		1000000,
		10000000,
		100000000,
	}
	// prefixCounts[i] means count of numbers share the same prefix
	// with tailing not greater than i
	// e.g. nums start with p, i = 1, count is 11 = p, p0~p9
	prefixCounts = []int{
		1,
		11,
		111,
		1111,
		11111,
		111111,
		1111111,
		11111111,
		111111111,
	}
)

// 0ms, 100%
// 1.9MB, 100%
func findKthNumber(n int, k int) int {
	prefix := 1
	k--
	tailWidth := len(strconv.Itoa(n)) - 1 // width of n - width of prefix
	for k > 0 && tailWidth > 0 {
		count := getCount(n, prefix, tailWidth)
		//fmt.Println(prefix, tailWidth, count, k)
		if k >= count {
			prefix++
			k -= count
			continue
		}

		prefix *= 10
		k--
		tailWidth--
	}
	//fmt.Println(prefix, k)

	if k > 0 {
		prefix += k
	}

	return prefix
}

// getCount of the nums with given prefix
func getCount(n int, prefix int, tailWidth int) int {
	//fmt.Println("getCount", prefix)
	// if prefix > n {
	// 	panic(prefix)
	// }
	if tailWidth == 0 {
		return 1
	}

	max := (prefix+1)*bases[tailWidth] - 1
	if max <= n {
		return prefixCounts[tailWidth]
	}

	count := getCount(n, prefix, tailWidth-1)
	if remain := n - prefix*bases[tailWidth] + 1; remain > 0 {
		count += remain
	}
	return count
}

func helper(num, end int) int {
	// 1st, try append 0
	if num*10 <= end {
		return num * 10
	}
	// 2nd, try + 1
	if num < end && num%10 < 9 {
		return num + 1
	}
	// 3rd, try remove tailing 9 and + 1
	num /= 10
	for num%10 == 9 {
		num /= 10
	}
	return num + 1
}

func main() {
	n := 3339
	x := 1
	for k := 1; k <= n; k++ {
		fmt.Print("#", k, ":", x, "-", findKthNumber(n, k), ",")
		x = helper(x, n)
	}
}
