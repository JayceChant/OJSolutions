package main

// 4ms, 90.24% (min 0ms)
// 2.8MB, 100%
func trap(height []int) int {
	sum := 0
	left := 0
	right := len(height) - 1
	for left < right {
		if height[left] < height[right] {
			h := height[left]
			left++
			for h >= height[left] {
				sum += (h - height[left])
				left++
			}
		} else {
			h := height[right]
			right--
			// maybe height[left] == height[right]
			for left < right && h >= height[right] {
				sum += (h - height[right])
				right--
			}
		}
	}
	return sum
}
