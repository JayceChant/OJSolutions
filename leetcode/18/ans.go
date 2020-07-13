package main

import "sort"

// 12ms (min 0ms)
// 2.7MB 100%
func fourSum1(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	size := len(nums)
	sort.Ints(nums)

	tt := target
	for i1 := 0; i1 < size-3; i1++ {
		n1 := nums[i1]
		if i1 > 0 && n1 == nums[i1-1] {
			continue
		}
		if n1 > 0 && n1 > tt {
			break
		}
		tt -= n1
		for i2 := i1 + 1; i2 < size-2; i2++ {
			n2 := nums[i2]
			if i2 > i1+1 && n2 == nums[i2-1] {
				continue
			}
			if n2 > 0 && n2 > tt {
				break
			}
			tt -= n2
			for i3 := i2 + 1; i3 < size-1; i3++ {
				n3 := nums[i3]
				if i3 > i2+1 && n3 == nums[i3-1] {
					continue
				}
				if (n3 > 0 && n3 > tt) || n3+nums[i3+1] > tt {
					break
				}
				tt -= n3
				i4 := size - 1
				for i4 > i3 && nums[i4] > tt {
					i4--
				}
				if i4 > i3 && nums[i4] == tt {
					ans = append(ans, []int{n1, n2, n3, nums[i4]})
				}
				tt += n3
			}
			tt += n2
		}
		tt += n1
	}

	return ans
}

// 12ms (min 0ms)
// 2.7MB 100%
func fourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	return nSum(4, nums, len(nums)-1, target, make([]int, 4), make([][]int, 0))
}

func nSum(n int, nums []int, end int, target int, cand []int, answers [][]int) [][]int {
	if n == 2 { // twoSum
		for right := end; right >= 1; right-- { // ← right
			num := nums[right]
			if right < end && num == nums[right+1] {
				continue // de-duplicate
			}
			if (num < 0 && num < target) || num+nums[right-1] < target {
				break // prune
			}
			target -= num
			left := 0
			for left < right && nums[left] < target { // left →
				left++
			}
			if left < right && nums[left] == target {
				ans := make([]int, len(cand))
				copy(ans, cand)
				ans[0] = target
				ans[1] = num
				answers = append(answers, ans)
			}
			target += num
		}
		return answers
	}
	n--
	for i := end; i >= n; i-- {
		if i < end && nums[i] == nums[i+1] {
			continue // de-duplicate
		}
		if nums[i] < 0 && nums[i] < target {
			break // prune
		}
		cand[n] = nums[i]
		answers = nSum(n, nums, i-1, target-nums[i], cand, answers)
	}
	return answers
}
