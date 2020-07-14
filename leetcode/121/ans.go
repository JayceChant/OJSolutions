package main

import "math"

func maxProfit(prices []int) int {
	maxPrfit := 0
	minPrice := math.MaxInt32
	for _, price := range prices {
		if minPrice > price {
			minPrice = price
		} else if profit := price - minPrice; maxPrfit < profit {
			maxPrfit = profit
		}
	}
	return maxPrfit
}
