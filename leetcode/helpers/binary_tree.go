package main

import (
	"fmt"
	"strconv"
	"strings"
)

// TreeNode ...
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func (tree *TreeNode) String() string {
	if tree == nil {
		return "nil"
	}

	left := tree.Left.String()
	right := tree.Right.String()
	return fmt.Sprintf("[%s,%d,%s]", left, tree.Val, right)
}

func main() {
	fmt.Println(listStrToBinaryTree("15,66,55,97,60,12,56,null,54,null,49,null,9,null,null,null,null,null,90"))
}

func listStrToBinaryTree(s string) *TreeNode {
	nodeStrs := strings.Split(s, ",")
	root := buildNode(nodeStrs[0])
	nodes := []*TreeNode{root}
	si := 1
	pi := 0
	for si < len(nodeStrs) {
		parent := nodes[pi]
		parent.Left = buildNode(nodeStrs[si])
		if parent.Left != nil {
			nodes = append(nodes, parent.Left)
		}
		si++
		if si == len(nodeStrs) {
			break
		}
		parent.Right = buildNode((nodeStrs[si]))
		if parent.Right != nil {
			nodes = append(nodes, parent.Right)
		}
		si++
		pi++
	}
	return root
}

func buildNode(s string) *TreeNode {
	s = strings.ToLower(strings.TrimSpace(s))
	if s == "null" || s == "nil" {
		return nil
	}
	val, _ := strconv.Atoi(s)
	return &TreeNode{Val: val}
}
