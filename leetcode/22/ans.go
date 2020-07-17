package main

// 0ms, 100%
// 2.6MB, 100%
func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	cand := make([]byte, n*2)
	return leftFirst(n, 0, 0, cand, ans)
}

func leftFirst(n int, left int, right int, cand []byte, ans []string) []string {
	//fmt.Println(left, right, string(cand[:left+right]))
	if left == n && right == n {
		return append(ans, string(cand))
	}

	for left < n {
		cand[left+right] = '('
		left++
		for r := right; r < left; {
			cand[left+r] = ')'
			r++
			ans = leftFirst(n, left, r, cand, ans)
		}
	}
	return ans
}

var (
	cache = map[int]([]string){
		0: []string{""},
	}
)

func generateParenthesisMemoDAC(n int) []string {
	if ans, ok := cache[n]; ok {
		return ans
	}

	ans := make([]string, 0)
	for ln := 0; ln < n; ln++ {
		for _, left := range generateParenthesisMemoDAC(ln) {
			for _, right := range generateParenthesisMemoDAC(n - ln - 1) {
				ans = append(ans, "("+left+")"+right)
			}
		}
	}
	cache[n] = ans
	return ans
}
