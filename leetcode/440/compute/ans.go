package main

import "fmt"

var (
	// base with same prefix
	prefixBases = []int{
		11111111,
		1111111,
		111111,
		11111,
		1111,
		111,
		11,
		1,
		0,
	}
	tenBases = []int{
		10000000,
		1000000,
		100000,
		10000,
		1000,
		100,
		10,
		1,
		0,
	}
)

func findKthNumber(n int, k int) int {
	ans := 0
	i := 0
	for n < tenBases[i] {
		i++
	}

	num := k - 1
	max := n
	ans = num/prefixBases[i] + 1 // 1st digit starts from 1
	num %= prefixBases[i]

	x := (max - prefixBases[i+1]*9) / tenBases[i] // x999.. <= max
	max -= (x + 1) * tenBases[i]                  // x9999... and x+1
	bound := x * prefixBases[i]
	//fmt.Println(x, max, bound)
	i++

	for j := i; max > 0 && tenBases[j] > 1; j++ {
		x := (max + 1) / tenBases[j]
		max -= x * tenBases[j]
		bound += 1 + x*prefixBases[j]
		//fmt.Println(x, max, bound)
	}

	if k <= bound {
		for num > 0 {
			num--
			ans = ans*10 + num/prefixBases[i]
			num %= prefixBases[i]
			i++
			//fmt.Println(ans)
		}
		return ans
	}

	k -= bound
	ans = n
	for k < prefixBases[i] {
		i++
	}
	for k > 0 {
		//fmt.Println(ans, k, i)
		if ans*10 <= n {
			ans *= 10
			k--
			continue
		}

		for k < prefixBases[i] {
			i++
		}
		if ans+tenBases[i] <= n && (ans/tenBases[i])%10 < 9 {
			ans += tenBases[i]
			k -= prefixBases[i]
			continue
		}

		ans /= 10
		for ans%10 == 9 {
			ans /= 10
		}
		ans++
		k--
	}
	return ans
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
	//findKthNumber(3339, 2603)
}

// func findKthNumber(n int, k int) int {
// 	ans := 0
// 	i := 0
// 	for n < prefixBases[i] {
// 		i++
// 	}

// 	bound := (((n + 1) / tenBases[i]) - 1) * prefixBases[i]
// 	if k <= bound {
// 		k--
// 		ans = k/prefixBases[i] + 1 // 1st digit starts from 1
// 		k %= prefixBases[i]
// 		for k > 0 {
// 			k--
// 			i++
// 			ans = ans*10 + k/prefixBases[i]
// 			k %= prefixBases[i]
// 		}
// 		return ans
// 	}

// 	return ans
// }
