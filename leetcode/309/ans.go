package main

// 给定价格求股市最大收益，只能全仓空仓，卖出后有一天冷冻期不能买入

func maxProfitN2(prices []int) int {
	// dp[i] means maxProfit of prices[i:]
	dp := make([]int, len(prices)+2)
	for buy := len(prices) - 2; buy >= 0; buy-- {
		dp[buy] = dp[buy+1]
		for sell := buy + 1; sell < len(prices); sell++ {
			profit := prices[sell] - prices[buy] + dp[sell+2]
			if dp[buy] < profit {
				dp[buy] = profit
			}
		}
	}
	return dp[0]
}

// 0 ms, faster than 100.00%
// 2.3 MB, less than 77.27% (min 2276kB, mine 2280kB)
func maxProfit(prices []int) int {
	n := len(prices)
	if n < 2 {
		return 0
	}
	// [empty]->(buy)->[full]->(sell)->[frozen]->(T+1)->[empty]
	// empty and full can be kept, but frozen not
	maxEmpty := 0
	maxFull := -prices[0]
	frozen := 0
	for i := 1; i < n; i++ {
		buyToday := maxEmpty - prices[i]
		unfreezeToday := frozen
		frozen = maxFull + prices[i] // sell today, unfreeze to empty next day

		if maxFull < buyToday {
			maxFull = buyToday // buy cheaper
		}
		if maxEmpty < unfreezeToday {
			maxEmpty = unfreezeToday
		}
	}
	if frozen > maxEmpty {
		return frozen
	}
	return maxEmpty
}
