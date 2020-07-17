package main

import "container/list"

// 28ms, 90.24%
// 6.1MB
func isBipartiteBFS(graph [][]int) bool {
	sets := make([]vertexSet, len(graph))
	queue := list.New()
	for i := 0; i < len(graph); i++ {
		if set := sets[i]; set != unknown {
			continue
		}

		sets[i] = setA
		queue.PushBack(i)
		for queue.Len() > 0 {
			e := queue.Front()
			queue.Remove(e)
			from := e.Value.(int)
			setFrom := sets[from]
			for _, to := range graph[from] {
				setTo := sets[to]
				if setTo == setFrom {
					return false
				}

				if setTo != unknown {
					continue
				}

				sets[to] = (setA + setB) - setFrom
				queue.PushBack(to)
			}
		}
	}
	return true
}

type vertexSet int8

const (
	unknown vertexSet = iota
	setA
	setB
)

var (
	sets  []vertexSet
	valid bool
)

// 28ms, 90.24%
// 6.1MB
func isBipartite(graph [][]int) bool {
	sets = make([]vertexSet, len(graph))
	valid = true
	for i := 0; valid && i < len(graph); i++ {
		if set := sets[i]; set != unknown {
			continue
		}

		isBipartiteDFS(graph, i, setA)
	}
	return valid
}

func isBipartiteDFS(graph [][]int, from int, setFrom vertexSet) {
	for _, to := range graph[from] {
		setTo := sets[to]
		if setTo == setFrom {
			valid = false
			return
		}

		if setTo != unknown {
			continue
		}

		sets[to] = (setA + setB) - setFrom
		isBipartiteDFS(graph, to, (setA+setB)-setFrom)
	}
}

// 36ms, 14.46%
func isBipartiteUF(graph [][]int) bool {
	uf := buildUnionFind(len(graph))
	for from := range graph {
		for i := range graph[from] {
			if uf.isUnited(from, graph[from][i]) {
				return false
			}
			if i > 0 {
				uf.union(graph[from][i-1], graph[from][i])
			}
		}
	}
	return true
}

type unionFind []int

func buildUnionFind(size int) unionFind {
	uf := make(unionFind, size)
	for i := 0; i < size; i++ {
		uf[i] = i
	}
	return uf
}

// with path shortedn 36ms
// without, 28ms
func (uf unionFind) findRoot(i int) int {
	// if uf[i] == i {
	// 	return i
	// }

	// uf[i] = uf.findRoot(uf[i])
	// return uf[i]
	p := uf[i]
	for p != uf[i] {
		i = p
		p = uf[i]
	}
	return p
}

func (uf unionFind) union(i, j int) {
	uf[uf.findRoot(j)] = uf.findRoot(i)
}

func (uf unionFind) isUnited(i, j int) bool {
	return uf.findRoot(i) == uf.findRoot(j)
}
