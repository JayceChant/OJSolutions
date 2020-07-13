package main

// 4ms (min 0ms)
// 3MB 100%
func intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) > len(nums2) {
		nums1, nums2 = nums2, nums1
	}
	numCount := make(map[int]int)
	for _, num := range nums1 {
		numCount[num]++
	}
	nums1 = nums1[:0]
	for _, num := range nums2 {
		if numCount[num] > 0 {
			nums1 = append(nums1, num)
			numCount[num]--
		}
	}
	return nums1
}
