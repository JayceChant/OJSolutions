package main

// 16ms (min 4ms)
func maxArea(height []int) int {
	left := 0
	right := len(height) - 1
	max := 0

	for left < right {
		//fmt.Println(left, right)
		area := min(height[left], height[right]) * (right - left)
		if max < area {
			max = area
		}
		if height[left] < height[right] {
			h := height[left]
			left++
			for /*left < right &&*/ h >= height[left] {
				left++
			}
		} else {
			h := height[right]
			right--
			for left < right && height[right] <= h {
				right--
			}
		}
	}
	return max
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
