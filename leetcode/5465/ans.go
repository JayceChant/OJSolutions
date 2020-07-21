package main

// 1 <= n <= 10^5
// edges.length == n - 1
// edges[i].length == 2
// 0 <= ai, bi < n
// ai != bi
// labels.length == n
// labels 仅由小写英文字母组成

var (
	// adj table
	tree    [][]int
	visited []bool
)

func countSubTrees(n int, edges [][]int, labels string) []int {
	tree = make([][]int, n)
	for i := 0; i < n; i++ {
		tree[i] = make([]int, 0)
	}
	visited = make([]bool, n)

	for _, edge := range edges {
		tree[edge[0]] = append(tree[edge[0]], edge[1])
		tree[edge[1]] = append(tree[edge[1]], edge[0])
	}

	ans := make([]int, n)
	count(0, labels, ans)
	return ans
}

func count(i int, labels string, ans []int) (letterCount []int) {
	visited[i] = true
	letterCount = make([]int, 26)
	for _, child := range tree[i] {
		if visited[child] {
			continue
		}
		lc := count(child, labels, ans)
		for i := range lc {
			letterCount[i] += lc[i]
		}
	}
	letterCount[labels[i]-'a']++
	ans[i] = letterCount[labels[i]-'a']
	return
}

// 4
// [[0,2],[0,3],[1,2]]
// "aeed"
// got [1,0,1,1]
// want [1,1,2,1]
