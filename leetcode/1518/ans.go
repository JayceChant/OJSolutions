package main

func numWaterBottles(numBottles int, numExchange int) int {
	if numExchange == 0 {
		return numBottles
	}

	sum := numBottles
	for numBottles >= numExchange {
		exchange := numBottles / numExchange
		sum += exchange
		numBottles = exchange + numBottles%numExchange
	}
	return sum
}
