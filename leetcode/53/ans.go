package main

import "math"

// 4ms, 96.71% (min 0ms)
// 3284, 54.55% (min 3280)
func maxSubArray(nums []int) int {
	dp := nums[0]
	max := dp
	for i := 1; i < len(nums); i++ {
		if dp > 0 {
			dp += nums[i]
		} else {
			dp = nums[i]
		}
		if max < dp {
			max = dp
		}
	}
	return max
}

// O(N) build tree, slower, take more space
// but every query take O(log(N)) every time after tree is built
func maxSubArrayST(nums []int) int {
	tree := buildTree(nums, 0, len(nums)-1)
	return tree.maxSum(0, len(nums)-1)
}

type sums struct {
	allSum   int
	leftSum  int
	rightSum int
	midSum   int
}

type treeNode struct {
	*sums
	start int
	end   int
	left  *treeNode
	right *treeNode
}

func (tree *treeNode) maxSum(start int, end int) int {
	sums := tree.maxSums(start, end)
	if sums != nil {
		return sums.midSum
	}
	return math.MinInt32
}

func (tree *treeNode) maxSums(start int, end int) *sums {
	if start > end || end < tree.start || start > tree.end {
		// no intersection
		return nil
	}
	if start <= tree.start && tree.end <= end {
		return tree.sums
	}
	return pushUp(tree.left.maxSums(start, end), tree.right.maxSums(start, end))
}

func buildTree(nums []int, start int, end int) *treeNode {
	node := &treeNode{
		start: start,
		end:   end,
	}

	// leaf
	if start == end {
		node.sums = &sums{
			allSum:   nums[start],
			leftSum:  nums[start],
			rightSum: nums[start],
			midSum:   nums[start],
		}
		return node
	}

	// build children
	mid := (start + end) / 2
	node.left = buildTree(nums, start, mid)
	node.right = buildTree(nums, mid+1, end)
	node.sums = pushUp(node.left.sums, node.right.sums)
	return node
}

func pushUp(sl, sr *sums) *sums {
	if sl == nil {
		return sr
	}
	if sr == nil {
		return sl
	}
	s := new(sums)
	s.allSum = sl.allSum + sr.allSum
	s.leftSum = max(sl.leftSum, sl.allSum+sr.leftSum)
	s.rightSum = max(sl.rightSum+sr.allSum, sr.rightSum)
	s.midSum = max(sl.midSum, sl.rightSum+sr.leftSum, sr.midSum)
	return s
}

func max(nums ...int) int {
	m := math.MinInt32
	for _, n := range nums {
		if m < n {
			m = n
		}
	}
	return m
}
