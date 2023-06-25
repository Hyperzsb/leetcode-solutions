func max(a, b int) int {
    if a > b {
        return a
    } else {
        return b
    }
}

func maxProfit(prices []int, fee int) int {
    buy, sell := math.MinInt32, 0

    for _, price := range prices {
        buy = max(buy, sell - price)
        sell = max(sell, buy + price - fee)
    }

    return sell
}

