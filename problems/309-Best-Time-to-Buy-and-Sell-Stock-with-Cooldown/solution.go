func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func maxProfit(prices []int) int {
	wait, hold, sell, preSell := 0, -prices[0], 0, 0
	for i := 1; i < len(prices); i++ {
		preSell = sell
		sell = hold + prices[i]
		hold = max(hold, wait-prices[i])
		wait = max(wait, preSell)
	}

	return max(wait, sell)
}

