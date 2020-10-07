package main

import "strconv"

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 4ms (min 0ms)
// 4.7MB
func recoverFromPreorder(S string) *TreeNode {
	str = S
	length = len(S)
	root, _, _ := rfp(0, 0)
	return root
}

var (
	str    string
	length int
)

// [st, ed)
func rfp(st int, depth int) (*TreeNode, int, int) {
	ed := st
	for ed < length && str[ed] != '-' {
		ed++
	}
	val, _ := strconv.Atoi(str[st:ed])
	root := &TreeNode{Val: val}
	st = ed

	for ed < length && str[ed] == '-' {
		ed++
	}
	nextDep := ed - st
	st = ed // skip '-'s

	if nextDep > depth {
		root.Left, st, nextDep = rfp(st, depth+1)
	}
	if nextDep > depth {
		root.Right, st, nextDep = rfp(st, depth+1)
	}
	return root, st, nextDep
}
