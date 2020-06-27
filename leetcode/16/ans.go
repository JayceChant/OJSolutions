package main

// 4ms, 98.48% (min 0ms)
// 2.7MB, 68.13% (mine 2724,min 2712kb)
func threeSumClosest(nums []int, target int) int {
	quickSort(nums)
	minDif := 100000
	ans := 0
	for i1 := 0; i1 < len(nums)-2; i1++ {
		i2 := i1 + 1
		i3 := len(nums) - 1
		for i2 < i3 {
			sum := nums[i1] + nums[i2] + nums[i3]
			if sum == target {
				return target
			}
			diff := target - sum
			if diff > 0 {
				i2++
				if diff < minDif {
					minDif = diff
					ans = sum
				}
			} else {
				i3--
				if -diff < minDif {
					minDif = -diff
					ans = sum
				}
			}
		}
	}
	return ans
}

func quickSort(nums []int) {
	partition(nums, 0, len(nums))
}

// [head, tail)
func partition(nums []int, head int, tail int) {
	if tail-head < 2 {
		return
	}

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
