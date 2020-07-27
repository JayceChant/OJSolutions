package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 28ms
// 7MB
func countPairs(root *TreeNode, distance int) int {
	count, _ := cp(root, distance)
	return count
}

// leaf[distance] = count
func cp(root *TreeNode, distance int) (count int, leafs [10]int) {
	if root.Left == nil && root.Right == nil {
		// leaf
		count = 0
		leafs[1] = 1
		return
	}

	var leftCount, rightCount int
	var leftLeaves, rightLeaves [10]int

	if root.Left != nil {
		leftCount, leftLeaves = cp(root.Left, distance)
		count += leftCount
	}

	if root.Right != nil {
		rightCount, rightLeaves = cp(root.Right, distance)
		count += rightCount
	}

	for lDis, lCnt := range leftLeaves {
		for rDis, rCnt := range rightLeaves {
			// depth sum <= distance
			if lDis+rDis <= distance {
				count += lCnt * rCnt
			}
		}
		lDis++
		if lDis < distance {
			leafs[lDis] = lCnt
		}
	}
	for rDis, rCnt := range rightLeaves {
		rDis++
		if rDis < distance {
			leafs[rDis] += rCnt
		}
	}
	return
}
