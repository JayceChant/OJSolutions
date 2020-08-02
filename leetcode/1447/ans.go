package main

import "fmt"

// 72ms
// 6.7MB
func simplifiedFractions(n int) (ans []string) {
	for ; n > 1; n-- {
		ans = append(ans, fmt.Sprintf("1/%d", n))
		if n > 2 {
			ans = append(ans, fmt.Sprintf("%d/%d", n-1, n))
		}
		for i := 2; i <= n/2; i++ {
			if isSimplified(i, n) {
				ans = append(ans, fmt.Sprintf("%d/%d", i, n), fmt.Sprintf("%d/%d", n-i, n))
			}
		}
	}
	return
}

func isSimplified(i, n int) bool {
	for i > 0 {
		n, i = i, n%i
	}
	return n == 1
}

// 80ms
// 6.6MB
func simplifiedFractionsHash(n int) (ans []string) {
	fracHash := map[float64]bool{0.5: true}
	if n >= 2 {
		ans = append(ans, "1/2")
	}
	for d := 3; d <= n; d++ {
		for i := 1; i <= d/2; i++ {
			val := float64(i) / float64(d)
			if !fracHash[val] {
				ans = append(ans, fmt.Sprintf("%d/%d", i, d), fmt.Sprintf("%d/%d", d-i, d))
				fracHash[val] = true
			}
		}
	}
	return
}
