package main

import (
	"sort"
	"strings"
)

// 16ms
// 5.7MB
func findItinerary1(tickets [][]string) []string {
	tt = tickets
	size = len(tickets)
	adj = make(map[string][]int)
	for i, t := range tickets {
		adj[t[0]] = append(adj[t[0]], i)
	}

	for _, tis := range adj {
		sort.Slice(tis, func(i, j int) bool {
			return strings.Compare(tickets[tis[i]][1], tickets[tis[j]][1]) < 0
		})
	}

	used = make([]bool, size)
	cand = make([]int, size)
	done = false
	dfs1(0, "JFK")
	ans := make([]string, 0, size+1)
	ans = append(ans, "JFK")
	for _, ti := range cand {
		ans = append(ans, tickets[ti][1])
	}
	return ans
}

var (
	tt   [][]string
	size int
	adj  map[string]([]int)
	used []bool
	cand []int
	done bool
)

func dfs1(i int, from string) {
	for _, ti := range adj[from] {
		if done {
			return
		}
		if used[ti] {
			continue
		}

		cand[i] = ti
		if i == size-1 {
			done = true
			return
		}

		used[ti] = true
		dfs1(i+1, tt[ti][1])
		used[ti] = false
	}
}

// 12ms
// 5.7MB
func findItinerary(tickets [][]string) []string {
	i := len(tickets)
	adj := make(map[string][]string)
	for _, t := range tickets {
		from, to := t[0], t[1]
		adj[from] = append(adj[from], to)
	}

	for _, tos := range adj {
		sort.Strings(tos)
	}

	ans := make([]string, i+1)

	var dfs func(string)
	dfs = func(from string) {
		for len(adj[from]) > 0 {
			tos := adj[from]
			to := tos[0]
			adj[from] = tos[1:]
			dfs(to)
		}
		ans[i] = from
		i--
	}

	dfs("JFK")

	return ans
}
