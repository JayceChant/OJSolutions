package main

// 2-loop:32ms; 1-loop-with-sum:28ms
// 6.3MB
func findErrorNumsHash(nums []int) []int {
	ans := make([]int, 2)
	size := len(nums)
	hash := make([]bool, size+1)
	sum := 0
	for _, n := range nums {
		sum += n
		if hash[n] {
			ans[0] = n
		} else {
			hash[n] = true
		}
	}
	ans[1] = ans[0] + size*(size+1)/2 - sum
	return ans
}

// 32ms
// 6.2MB
func findErrorNums(nums []int) []int {
	xor := 0
	sum := 0
	size := len(nums)
	for i, n := range nums {
		sum += n
		xor = xor ^ (i + 1) ^ n
	}
	lowbit := xor & (-xor)
	xor0 := 0
	xor1 := 0
	for i, n := range nums {
		id := i + 1
		if n&lowbit == 0 {
			xor0 ^= n
		} else {
			xor1 ^= n
		}
		if id&lowbit == 0 {
			xor0 ^= id
		} else {
			xor1 ^= id
		}
	}
	if (xor0 < xor1) == (sum < size*(size+1)/2) {
		return []int{xor0, xor1}
	}
	return []int{xor1, xor0}
}
