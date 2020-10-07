package main

// TreeNode ..
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 4ms
// 5.5MB
func isValidBSTRec(root *TreeNode) bool {
	if root == nil {
		return true
	}
	valid, _ := validate(root, nil)
	return valid
}

func validate(root *TreeNode, prev *TreeNode) (bool, *TreeNode) {
	var valid bool
	if root.Left != nil {
		valid, prev = validate(root.Left, prev)
		if !valid {
			return false, nil
		}
	}

	if prev != nil && prev.Val >= root.Val {
		return false, nil
	}
	prev = root

	if root.Right != nil {
		valid, prev = validate(root.Right, prev)
		if !valid {
			return false, nil
		}
	}
	return true, prev
}

// Morris
// 8ms
// 5.4MB
func isValidBST(root *TreeNode) bool {
	valid := true
	var pred, prev *TreeNode
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
				if valid && prev != nil {
					valid = prev.Val < root.Val
				}
				// disconnect
				pred.Right = nil
				// root.Left is finished
				prev = root
				root = root.Right
			}
		} else {
			if valid && prev != nil {
				valid = prev.Val < root.Val
			}
			prev = root
			root = root.Right
		}
	}
	return valid
}
