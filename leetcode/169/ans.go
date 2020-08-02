package main

import (
	"math/rand"
	"sort"
)

// 24ms
// 6MB
func majorityElementHashCount(nums []int) int {
	threshold := len(nums) / 2
	counter := map[int]int{}
	for _, num := range nums {
		counter[num]++
		if counter[num] > threshold {
			return num
		}
	}
	return 0
}

// 24ms
// 5.9MB
func majorityElementDAC(nums []int) int {
	num, _ := partition(nums, len(nums))
	return num
}

func partition(nums []int, size int) (int, int) {
	newSize := size & (size - 1)
	var num1, diff1 int
	if newSize > 0 {
		num1, diff1 = partition(nums, newSize)
	}
	num2, diff2 := countNum(nums, newSize, size)
	if num1 == num2 {
		return num1, diff1 + diff2
	}

	if diff1 > diff2 {
		return num1, diff1 - diff2
	}

	if diff1 < diff2 {
		return num2, diff2 - diff1
	}
	return 0, 0
}

// [start, end)
// end - start must be 2^n
func countNum(nums []int, start, end int) (int, int) {
	if end-start == 1 {
		return nums[start], 1
	}

	mid := (start + end) / 2
	num1, diff1 := countNum(nums, start, mid)
	num2, diff2 := countNum(nums, mid, end)
	if num1 == num2 {
		return num1, diff1 + diff2
	}

	if diff1 > diff2 {
		return num1, diff1 - diff2
	}

	if diff1 < diff2 {
		return num2, diff2 - diff1
	}
	return 0, 0
}

// 24ms
// 5.9MB
func majorityElementSort(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 20ms
// 5.9MB
func majorityElementBM(nums []int) int {
	cand := 0
	count := 0
	size := len(nums)
	half := 0
	for i := 0; i < size; i++ {
		if count == 0 {
			cand = nums[i]
			half = (size - i) / 2
		}

		if nums[i] == cand {
			count++
			if count > half {
				return cand
			}
		} else {
			count--
		}
	}
	return cand
}

// 16ms
// 6MB
func majorityElementRand(nums []int) int {
	size := len(nums)
	half := size / 2
	for {
		cand := nums[rand.Intn(size)]
		count := 0
		for i := 0; i < size; i++ {
			if nums[i] == cand {
				count++
				if count > half {
					return cand
				}
			}
			// // time cost more stable after add break code,
			// // make the fastest cases slower(more judgement), and the slowest faster(early break)
			// else if i-count > half {
			// 	break
			// }
		}
	}
}
