package main

// 44ms
// 6.6MB
func smallestDivisor(nums []int, threshold int) int {
	sum := 0
	high := 1
	for _, num := range nums {
		sum += num
		if high < num {
			high = num
		}
	}
	if sum <= threshold {
		return 1
	}
	low := sum / threshold
	high = ceilingDiv(high, threshold/len(nums))
	for low < high {
		div := (low + high) / 2
		sum = numsDiv(nums, div)
		if sum <= threshold {
			high = div
		} else {
			low = div + 1
		}
	}
	return high
}

func ceilingDiv(num, div int) int {
	return (num + div - 1) / div
}

func numsDiv(nums []int, div int) int {
	sum := 0
	for _, num := range nums {
		sum += ceilingDiv(num, div)
	}
	return sum
}
