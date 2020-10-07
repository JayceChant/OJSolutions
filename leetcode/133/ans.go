package main

// Node ...
type Node struct {
	Val       int
	Neighbors []*Node
}

// 0ms
// 2.7MB
func cloneGraph(node *Node) *Node {
	if node == nil {
		return nil
	}
	cloneMap = make(map[*Node]*Node)
	return cloneNode(node)
}

var (
	cloneMap map[*Node]*Node
)

func cloneNode(node *Node) *Node {
	clone := &Node{
		Val:       node.Val,
		Neighbors: make([]*Node, len(node.Neighbors)),
	}
	cloneMap[node] = clone
	for i, nb := range node.Neighbors {
		if nbc, ok := cloneMap[nb]; ok {
			clone.Neighbors[i] = nbc
			continue
		}

		clone.Neighbors[i] = cloneNode(nb)
	}
	return clone
}
