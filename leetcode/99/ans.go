package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// no pruning
// 12ms
// 6MB
// with done pruning
// 8ms
// 6MB
func recoverTreeRec(root *TreeNode) {
	done = false
	newPair = false
	inverse = make([]*TreeNode, 0)
	findInverse(root, nil)
	last := len(inverse) - 1
	inverse[0].Val, inverse[last].Val = inverse[last].Val, inverse[0].Val
}

var (
	done    bool
	newPair bool
	inverse []*TreeNode
)

func findInverse(root *TreeNode, prev *TreeNode) *TreeNode {
	if root.Left != nil {
		prev = findInverse(root.Left, prev)
		if newPair {
			if inverse[0].Val < root.Val {
				done = true
			}
			newPair = false
		}
	}

	if !done && prev != nil && prev.Val > root.Val {
		inverse = append(inverse, prev, root)
		if len(inverse) == 4 {
			done = true
		} else {
			newPair = true
		}
	}
	prev = root

	if !done && root.Right != nil {
		prev = findInverse(root.Right, prev)
		if newPair {
			if inverse[0].Val < root.Val {
				done = true
			}
			newPair = false
		}
	}

	return prev
}

// Pointing the right pointer of the preceding node to the root,
// thus avoiding the use of the (call) stack,
// reduces the extra space to a constant.
// But finding the predecessor node,
// as well as detecting the loop to determine the return to the root and disconnect it,
// requires additional traversal.
// It is a trade-off of time for space.
// no pruning
// 16ms
// 6MB
// a != nil pruning?
// no way. need to iterate over the remaining nodes
// to disconnect temp link, no pruning
func recoverTree(root *TreeNode) {
	var a, b, prev, pred *TreeNode
	for root != nil {
		if root.Left != nil {
			pred = root.Left
			for pred.Right != nil && pred.Right != root {
				pred = pred.Right
			}

			if pred.Right == nil {
				pred.Right = root
				root = root.Left
			} else { // pred.Right == root
				if prev != nil && prev.Val > root.Val {
					b = root
					if a == nil {
						a = prev
					}
				}

				prev = root
				pred.Right = nil
				root = root.Right
			}
		} else { // root.Left == nil
			if prev != nil && prev.Val > root.Val {
				b = root
				if a == nil {
					a = prev
				}
			}
			prev = root
			root = root.Right
		}
	}
	a.Val, b.Val = b.Val, a.Val
}
