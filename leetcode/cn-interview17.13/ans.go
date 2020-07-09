package main

func respace(dictionary []string, sentence string) int {
	root := new(treeNode)
	for _, word := range dictionary {
		root.addWord(word)
	}
	n := len(sentence)
	dp := make([]int, n+1)
	// [start, end)
	for start := n - 1; start >= 0; start-- {
		node := root
		dp[start] = dp[start+1] + 1
		for i := start; i < n; {
			idx := sentence[i] - 'a'
			node = node.next[idx]
			if node == nil {
				break
			}
			i++
			if node.terminal && dp[start] > dp[i] {
				dp[start] = dp[i]
			}
		}
	}
	// fmt.Println(dp)
	return dp[0]
}

type treeNode struct {
	// 0-a, 1-b, 2-c ... 25-z
	next     [26]*treeNode
	terminal bool
}

func (node *treeNode) addWord(word string) {
	for i := range word {
		idx := word[i] - 'a'
		if node.next[idx] == nil {
			node.next[idx] = new(treeNode)
		}
		node = node.next[idx]
	}
	node.terminal = true
}
