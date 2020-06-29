package main

import (
	"math/rand"
)

// 32 ms, faster than 73.82%
// 7 MB, less than 48.09%
func threeSum(nums []int) [][]int {
	n := len(nums)
	answers := make([][]int, 0)
	if n <= 2 {
		return answers
	}
	quickSort(nums)
	for i1 := 0; i1 < n-2; i1++ {
		target := -nums[i1]
		if target < 0 {
			break
		}
		if i1 > 0 && nums[i1-1] == nums[i1] {
			continue
		}
		i2 := i1 + 1
		i3 := n - 1
		for i2 < i3 {
			sum := nums[i2] + nums[i3]
			if sum == target {
				answers = append(answers, []int{nums[i1], nums[i2], nums[i3]})
				i2++
				for i2 < i3 && nums[i2-1] == nums[i2] {
					i2++
				}
				i3--
				for i2 < i3 && nums[i3] == nums[i3+1] {
					i3--
				}
			} else if sum < target {
				i2++
			} else { // sum > target
				i3--
			}
		}
	}
	return answers
}

func quickSort(nums []int) {
	partition(nums, 0, len(nums))
}

// [head, tail)
func partition(nums []int, head int, tail int) {
	if tail-head < 2 {
		return
	}
	x := rand.Intn(tail - head)
	nums[head], nums[head+x] = nums[head+x], nums[head]

	left := head + 1
	right := tail - 1
	for {
		for left < tail && nums[left] <= nums[head] {
			left++
		}
		for nums[right] > nums[head] {
			right--
		}
		if left >= right {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
	}
	nums[head], nums[right] = nums[right], nums[head]
	partition(nums, head, right)
	partition(nums, right+1, tail)
}

// cn :228 ms, 在所有 Go 提交中击败了 14.37%
// 6.8 MB, 在所有 Go 提交中击败了 93.33%
func threeSumMap(nums []int) [][]int {
	answers := make([][]int, 0)
	numToCount := make(map[int]uint8)
	for _, num := range nums {
		numToCount[num]++
	}
	for num1, count1 := range numToCount {
		if num1 == 0 {
			if count1 >= 3 {
				answers = append(answers, []int{0, 0, 0})
			}
		} else if count1 >= 2 && numToCount[-2*num1] > 0 {
			if num1 <= 0 {
				answers = append(answers, []int{num1, num1, -2 * num1})
			} else {
				answers = append(answers, []int{-2 * num1, num1, num1})
			}
		}
		for num2 := range numToCount {
			if num2 <= num1 {
				continue
			}
			num3 := -(num1 + num2)
			if num3 <= num2 {
				continue
			}
			if numToCount[num3] > 0 {
				answers = append(answers, []int{num1, num2, num3})
			}
		}
	}
	return answers
}
