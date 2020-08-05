package main

// 12ms
// 6MB
func canFinish(numCourses int, prerequisites [][]int) bool {
	adj = make([][]int, numCourses)
	for i := 0; i < numCourses; i++ {
		adj[i] = make([]int, 0)
	}

	for _, pre := range prerequisites {
		adj[pre[1]] = append(adj[pre[1]], pre[0])
	}

	states = make([]state, numCourses)

	for v := 0; v < numCourses; v++ {
		if states[v] != open {
			continue
		}

		states[v] = inPath
		if !dfs(v) {
			return false
		}
		states[v] = visited
	}
	return true
}

type state int8

const (
	open state = iota
	inPath
	visited
)

var (
	adj    [][]int
	states []state
)

func dfs(v int) bool {
	for _, w := range adj[v] {
		if states[w] == inPath {
			return false
		}

		if states[w] == visited {
			continue
		}

		states[w] = inPath
		if !dfs(w) {
			return false
		}
		states[w] = visited
	}
	return true
}
