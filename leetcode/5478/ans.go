package main

func maxSum(nums1 []int, nums2 []int) int {
	size1 := len(nums1)
	size2 := len(nums2)
	var maxSum1, maxSum2 uint64
	var i1, i2 int
	for i1 < size1 && i2 < size2 {
		for i1 < size1 && nums1[i1] < nums2[i2] {
			maxSum1 += uint64(nums1[i1])
			i1++
		}

		if i1 == size1 {
			break
		}

		for i2 < size2 && nums2[i2] < nums1[i1] {
			maxSum2 += uint64(nums2[i2])
			i2++
		}

		if i2 == size2 {
			break
		}

		//fmt.Println(nums1[i1], nums2[i2])
		if nums1[i1] != nums2[i2] {
			continue
		}

		maxSum1 += uint64(nums1[i1])
		i1++
		maxSum2 += uint64(nums2[i2])
		i2++

		// sum to a same num
		if maxSum1 < maxSum2 {
			maxSum1 = maxSum2
		} else {
			maxSum2 = maxSum1
		}
	}

	for i1 < size1 {
		maxSum1 += uint64(nums1[i1])
		i1++
	}

	for i2 < size2 {
		maxSum2 += uint64(nums2[i2])
		i2++
	}

	if maxSum1 < maxSum2 {
		return int(maxSum2 % modBase)
	}
	return int(maxSum1 % modBase)
}

const (
	modBase = 1000000007
)

// func sumHelper(nums []int) uint64 {
// 	var sum uint64
// 	for _, num := range nums {
// 		sum += uint64(num)
// 	}
// 	return sum
// }

// func findSameHelper(nums1 []int, nums2 []int) []int {
// 	size1 := len(nums1)
// 	size2 := len(nums2)
// 	ans := make([]int, 0)
// 	for i1 := 0; i1 < size1; i1++ {
// 		for i2 := 0; i2 < size2; i2++ {
// 			if nums2[i2] == nums1[i1] {
// 				ans = append(ans, nums2[i2])
// 				break
// 			}
// 			if nums2[i2] > nums1[i1] {
// 				break
// 			}
// 		}
// 	}
// 	return ans
// }
