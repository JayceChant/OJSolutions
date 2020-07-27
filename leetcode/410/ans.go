package main

// 0ms, 100%
// 2.1MB, 54.55%
func splitArray(nums []int, m int) int {
	size = len(nums)
	low := nums[0]

	sums = make([]int, size)
	sums[0] = nums[0]
	for i := 1; i < size; i++ {
		if low < nums[i] {
			low = nums[i]
		}
		sums[i] = sums[i-1] + nums[i]
	}

	if size == m {
		return low
	}

	if m == 1 {
		return sums[size-1]
	}

	width := size / m
	if size%m != 0 {
		width++
	}

	//fmt.Println(sums)
	//fmt.Println("width", width)

	high := sums[width-1]
	for i := width; i < size; i++ {
		sum := sums[i] - sums[i-width]
		if high < sum {
			high = sum
		}
	}

	for low < high {
		mid := (low + high) / 2
		//fmt.Println(low, mid, high)

		if checkSum(mid, m) {
			high = mid
		} else {
			low = mid + 1
		}
	}
	return high
}

var (
	size int
	sums []int
)

func checkSum(sum int, count int) bool {
	lastSum := 0
	i := 0
	for count > 1 && i < size {
		for i < size && sums[i]-lastSum <= sum {
			i++
		}
		lastSum = sums[i-1]
		count--
	}
	return i == size || sums[size-1]-lastSum <= sum
}

// TODO: try DP
