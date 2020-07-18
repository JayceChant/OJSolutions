package main

import (
	"math"
)

// TLE
func splitArrayN2(nums []int) int {
	size := len(nums)
	if calcGCD(nums[0], nums[size-1]) > 1 {
		return 1
	}
	// dp[i] means splitArray(nums[i:])
	dp := make([]int, size+1)
	// 0 <= left < right <= size-1
	for left := size - 1; left >= 0; left-- {
		minRight := dp[left+1]
		dp[left] = minRight + 1
		for right := size - 1; right > left; right-- {
			if dp[right+1] >= minRight {
				continue
			}

			if calcGCD(nums[left], nums[right]) > 1 {
				minRight = dp[right+1]
				dp[left] = minRight + 1
			}
		}
	}
	//fmt.Println(dp)
	return dp[0]
}

func calcGCD(a, b int) int {
	// make sure a >= b
	if a < b {
		a, b = b, a
	}

	for b > 0 {
		a, b = b, a%b
	}
	return a
}

// =======================================TLE code above=============================================

func splitArray(nums []int) int {
	primeToMin := map[int]int{}
	current := 0
	calcMinPrimeFactors()
	for _, num := range nums {
		current++
		// make a copy of last current + 1, to update primeToMin
		// avoiding the reduction from primeToMin to current reduce another prime record again
		// e.g. 2 -> 6 -> 3
		lastAdd1 := current
		pfs := getPrimeFactors(num)
		//fmt.Println("pfs:", pfs)
		for _, pf := range pfs {
			min, exist := primeToMin[pf]
			if exist && current > min {
				current = min
			}

			// use lastAdd1 here to avoid the reduction of current spread to other prime factors
			if !exist || min > lastAdd1 {
				primeToMin[pf] = lastAdd1
			}
		}
		//fmt.Println(num, current, primeToMin)
	}
	return current
}

func getPrimeFactors(num int) []int {
	res := make([]int, 0)

	pf := 0
	for num > 1 {
		if pf != minPrimeFactors[num] {
			pf = minPrimeFactors[num]
			res = append(res, pf)
		}
		num /= pf
	}

	return res
}

const (
	maxNum = 1000000
)

var (
	minPrimeFactors [maxNum + 1]int
	done            = false
)

// 1548520 ns/op for 1000000, 20x faster
func calcMinPrimeFactors() {
	if done {
		return
	}
	done = true
	p := 2
	for p <= maxNum {
		minPrimeFactors[p] = p

		num := p * p
		for num <= maxNum {
			minPrimeFactors[num] = p
			num += p
		}
		p++

		for p <= maxNum {
			if minPrimeFactors[p] == 0 {
				break
			}
			p++
		}
	}
}

// =============================TLE code below==========================================

var minPrimeFactors2 [maxNum + 1]int

// TLE
// 37582797 ns/op for 1000000, 20x slower
func calcMinPrimeFactors2() {
	// if done {
	// 	return
	// }
	//done = true
	primeNums := []int{2, 3, 5, 7}
	for num := 3; num <= maxNum; num += 2 {
		minPrimeFactors2[num-1] = 2
		sqrt := int(math.Sqrt(float64(num)))
		for _, pn := range primeNums {
			if pn > sqrt {
				break
			}

			if num%pn == 0 {
				minPrimeFactors2[num] = pn
			}
		}

		if minPrimeFactors2[num] == 0 {
			primeNums = append(primeNums, num)
			minPrimeFactors2[num] = num
		}
	}
}

// func main() {
// 	calcMinPrimeFactors()
// 	fmt.Println(minPrimeFactors)
// }
