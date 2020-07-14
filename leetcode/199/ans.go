package main

import "fmt"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 0ms, 100%
// 2.2MB, 100%
func rightSideView(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	ans, _ = rightSideViewRec(root, 0, 0, ans)
	return ans
}

func rightSideViewRec(root *TreeNode, depth int, maxDepth int, ans []int) ([]int, int) {
	depth++
	fmt.Println(root.Val, depth, maxDepth)

	if depth > maxDepth {
		// visible only when current depth > max depth of right trees
		ans = append(ans, root.Val)
		maxDepth = depth
	}

	var childDepth int
	if root.Right != nil {
		ans, childDepth = rightSideViewRec(root.Right, depth, maxDepth, ans)
		if maxDepth < childDepth {
			maxDepth = childDepth
		}
	}
	if root.Left != nil {
		ans, childDepth = rightSideViewRec(root.Left, depth, maxDepth, ans)
		if maxDepth < childDepth {
			maxDepth = childDepth
		}
	}
	return ans, maxDepth
}

// 0ms, 100%
// 2.3MB, 100%
func rightSideViewBFS(root *TreeNode) []int {
	ans := make([]int, 0)
	if root == nil {
		return ans
	}

	queue := deque{}
	queue.init()
	queue.push(&queNode{
		treeNode: root,
		depth:    1,
	})

	pool := deque{}
	pool.init()

	depth := 0
	for queue.size > 0 {
		entry := queue.pop()
		root = entry.treeNode
		if depth < entry.depth {
			ans = append(ans, root.Val)
			depth++
		}

		// reuse entry
		if root.Right != nil {
			entry.treeNode = root.Right
			entry.depth = depth + 1
			queue.push(entry)
			entry = nil
		}

		if root.Left != nil {
			if entry == nil {
				if pool.size > 0 {
					entry = pool.pop()
				} else {
					entry = new(queNode)
				}
			}
			entry.treeNode = root.Left
			entry.depth = depth + 1
			queue.push(entry)
			entry = nil
		}

		if entry != nil {
			pool.push(entry)
		}
	}
	return ans
}

type queNode struct {
	treeNode *TreeNode
	depth    int
	prev     *queNode
	next     *queNode
}

type deque struct {
	root queNode
	size int
}

func (q *deque) init() {
	sentry := &q.root
	q.root.prev = sentry // tail
	q.root.next = sentry // head
}

func (q *deque) push(node *queNode) {
	tail := q.root.prev
	node.prev = tail
	node.next = tail.next
	tail.next = node
	q.root.prev = node
	q.size++
}

func (q *deque) pop() *queNode {
	head := q.root.next
	head.prev.next = head.next
	head.next.prev = head.prev
	// avoid memory leaks
	head.prev = nil
	head.next = nil
	q.size--
	return head
}
