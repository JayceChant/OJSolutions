package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// [10]int counter
// 184ms
// 21.4MB
// bitmap
// 148ms (min 120ms)
// 17MB
func pseudoPalindromicPaths(root *TreeNode) int {
	if root == nil {
		return 0
	}
	numBit = 0
	count = 0
	ppp(root)
	return count
}

var (
	// the (i)th bit is 1 means the num of i is odd
	// e.g. 10100 means the num of 2 and 4 is odd, and other nums are even
	numBit int
	count  int
)

func ppp(root *TreeNode) {
	numBit ^= 1 << root.Val

	if root.Left == nil && root.Right == nil {
		// leaf
		if numBit == 0 || numBit&(numBit-1) == 0 {
			// there's at most one 1 in numBit
			count++
		}
		numBit ^= 1 << root.Val
		return
	}

	if root.Left != nil {
		ppp(root.Left)
	}
	if root.Right != nil {
		ppp(root.Right)
	}

	numBit ^= 1 << root.Val
}
