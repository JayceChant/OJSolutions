package main

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 12ms (100%)
// 6.7MB
func allPossibleFBT(N int) []*TreeNode {
	if N&1 == 0 {
		return []*TreeNode{}
	}

	if trees, ok := cache[N]; ok {
		return trees
	}

	N--
	trees := make([]*TreeNode, 0)
	for left := 1; left < N; left += 2 {
		ltrees := allPossibleFBT(left)
		rtrees := allPossibleFBT(N - left)
		for _, lt := range ltrees {
			for _, rt := range rtrees {
				trees = append(trees, &TreeNode{
					Left:  lt,
					Right: rt,
				})
			}
		}
	}
	cache[N+1] = trees
	return trees
}

var (
	cache = map[int][]*TreeNode{
		1: []*TreeNode{new(TreeNode)},
	}
)
