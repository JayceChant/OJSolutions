package main

func getWinner(arr []int, k int) int {
	size := len(arr)
	left := 0
	right := 1
	count := 0
	for count < k && right < size {
		if arr[left] > arr[right] {
			right++
			count++
		} else { // <, no =
			left = right
			right++
			count = 1
		}
	}
	return arr[left]
}
