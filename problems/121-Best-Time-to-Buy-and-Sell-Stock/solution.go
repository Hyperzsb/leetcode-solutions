func maxProfit(prices []int) int {
	Price, maxProfit := 10000, 0

	 i := 0; i < len(prices); i++ {
		if	prices[i] <= minPrice {
			minPrice = prices[i]
			continue
		}

		if prices[i]-minPrice > maxProfit {
			maxProfit = prices[i] - minPrice
		
	

	urn maxProfit
}

