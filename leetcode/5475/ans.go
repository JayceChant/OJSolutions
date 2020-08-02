package main

// 0 <= i < j < k < arr.length
// |arr[i] - arr[j]| <= a
// |arr[j] - arr[k]| <= b
// |arr[i] - arr[k]| <= c
func countGoodTriplets(arr []int, a int, b int, c int) (count int) {
	size := len(arr)
	for i := 0; i < size; i++ {
		for j := i + 1; j < size; j++ {
			if myAbs(arr[i]-arr[j]) > a {
				continue
			}
			for k := j + 1; k < size; k++ {
				if myAbs(arr[j]-arr[k]) > b || myAbs(arr[i]-arr[k]) > c {
					continue
				}
				count++
			}
		}
	}
	return
}

func myAbs(num int) int {
	if num >= 0 {
		return num
	}
	return -num
}
