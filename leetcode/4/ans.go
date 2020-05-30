package main

const (
	intMax = int(^uint(0) >> 1)
	intMin = ^intMax
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	size := m + n
	half := size / 2

	if m > n {
		// 保持 nums1 大小 <= nums2 大小
		// 避免 pos 越界判断
		nums1, m, nums2, n = nums2, n, nums1, m
	}

	// left bound, right bound
	lb, rb := 0, m
	for {
		// pr = pointer right分界线右边的下标，范围[0, m]
		// 取 0 时左子集为空，取 m 时右子集为空
		pr1 := (lb + rb) / 2
		// left max of nums1
		lmax1 := intMin
		if pr1 > 0 {
			lmax1 = nums1[pr1-1]
		}
		// right min of nums1
		rmin1 := intMax
		if pr1 < m {
			rmin1 = nums1[pr1]
		}

		pr2 := half - pr1
		// left max of nums2
		lmax2 := intMin
		if pr2 > 0 {
			lmax2 = nums2[pr2-1]
		}
		// right min of nums2
		rmin2 := intMax
		if pr2 < n {
			rmin2 = nums2[pr2]
		}

		if rmin2 < lmax1 {
			rb = pr1
			continue
		}

		if rmin1 < lmax2 {
			lb = pr1 + 1
			continue
		}

		if lmax2 > lmax1 {
			lmax1 = lmax2
		}

		if rmin2 < rmin1 {
			rmin1 = rmin2
		}

		if size%2 == 1 {
			return float64(rmin1)
		}

		return float64(lmax1+rmin1) / 2
	}
}

/** samples for step-by-step explanation */

func sample1(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	size := m + n
	// You may assume nums1 and nums2 cannot be both empty.
	// if size == 0 {
	// 	// edge case
	// 	// 这里不处理，后面 k-1 下标会出错
	// 	return 0.0
	// }

	nums := make([]int, 0, size)
	i1, i2 := 0, 0
	// 归并排序
	for i1 < m && i2 < n {
		if nums1[i1] < nums2[i2] {
			nums = append(nums, nums1[i1])
			i1++
		} else {
			nums = append(nums, nums2[i2])
			i2++
		}
	}
	// 如果其中一个数组有剩余元素，全部追加到最后
	if i1 < m {
		nums = append(nums, nums1[i1:]...)
	} else {
		nums = append(nums, nums2[i2:]...)
	}

	if size%2 == 1 {
		// 奇数个元素
		return float64(nums[size/2])
	}
	// 偶数个元素
	half := size / 2
	return float64(nums[half-1]+nums[half]) / 2
}

func sample2(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)
	size := m + n
	// 相当于 floor(size/2) ，用作下标相当于第 k+1 个元素
	// 奇数个元素时是正好是中位数，偶数时是较大的那个
	half := size / 2
	i1, i2 := 0, 0
	mleft, mright := 0, 0
	for i1+i2 <= half {
		// 有最外层的条件限制， i1, i2 不可能同时到达尾部
		// 所以一旦其中一个数组到最后，另一个数组无需比较继续往后遍历
		if i1 < m && (i2 == n || nums1[i1] < nums2[i2]) {
			mleft = mright
			mright = nums1[i1]
			i1++
		} else {
			mleft = mright
			mright = nums2[i2]
			i2++
		}
	}

	if size%2 == 1 {
		return float64(mright)
	}
	return float64(mleft+mright) / 2
}
