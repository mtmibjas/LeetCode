func maxProfit(prices []int) int {
	min := prices[0]
	profit := 0

	for i := 1; i < len(prices); i++ {
		if min > prices[i] {
			min = prices[i]
		} else if profit < (prices[i] - min) {
			profit = prices[i] - min
		}

	}
	return profit
}
