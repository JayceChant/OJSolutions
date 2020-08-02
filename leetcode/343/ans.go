package main

// 0ms
// 2.1MB
func integerBreak(n int) int {
	i := len(ans) / 2
	for len(ans) <= n {
		ans = append(ans, max(ans[i]*ans[i], ans[i-1]*ans[i+1]), max(ans[i]*ans[i+1], ans[i-1]*ans[i+2]))
		i++
	}
	return ans[n]
}

var (
	ans = []int{
		0,  // 0
		0,  //1
		1,  //2
		2,  //3
		4,  //4
		6,  //5
		9,  //6
		12, //7
		18, //8
		27, //9
	}
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 0ms
// 1.9MB
func integerBreakMath(n int) int {
	if n <= 3 {
		return n - 1
	}

	switch n % 3 {
	case 0:
		return quickPow3(n / 3)
	case 1:
		return quickPow3(n/3-1) * 4
	default:
		return quickPow3(n/3) * 2
	}
}

func quickPow3(n int) int {
	res := 1
	d := 3
	for n > 0 {
		if n&1 == 1 {
			res *= d
		}
		n >>= 1
		d *= d
	}
	return res
}
