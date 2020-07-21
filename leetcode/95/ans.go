package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 4ms, 78.4%
func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return []*TreeNode{}
	}
	return gen(1, n)
}

var (
	// with no cache: 4.4 MB
	// with cache: 4MB, 100%
	cache    = map[int]([]*TreeNode){}
	nilTrees = []*TreeNode{nil}
)

func gen(st, ed int) []*TreeNode {
	key := st*10 + ed
	if trees, exist := cache[key]; exist {
		return trees
	}

	if st == ed {
		trees := []*TreeNode{&TreeNode{Val: st}}
		cache[key] = trees
		return trees
	}

	trees := make([]*TreeNode, 0)
	for root := st; root <= ed; root++ {
		var lefts, rights []*TreeNode
		if st < root {
			lefts = gen(st, root-1)
		} else {
			lefts = nilTrees
		}

		if root < ed {
			rights = gen(root+1, ed)
		} else {
			rights = nilTrees
		}

		for _, left := range lefts {
			for _, right := range rights {
				trees = append(trees, &TreeNode{
					Val:   root,
					Left:  left,
					Right: right,
				})
			}
		}
	}
	cache[key] = trees
	return trees
}
