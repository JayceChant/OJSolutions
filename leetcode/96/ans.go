package main

var (
	memo = map[int]int{}
)

func numTrees(n int) int {
	if n <= 1 {
		return 1
	}
	if ans, ok := memo[n]; ok {
		return ans
	}

	sum := 0
	for root := 1; root <= n; root++ {
		sum += numTrees(root-1) * numTrees(n-root)
	}
	memo[n] = sum
	return sum
}

// try catalan numbers
