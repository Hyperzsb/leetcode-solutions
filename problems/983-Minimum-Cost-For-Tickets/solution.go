func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func mincostTickets(days []int, costs []int) int {
	dp := make([]int, days[len(days)-1]+1)
	dp[0] = 0

	idx := 0
	for i := 1; i < len(dp); i++ {
		for i < days[idx] {
			dp[i] = dp[i-1]
			i++
		}
		idx++

		dp[i] = dp[i-1] + costs[0]
		dp[i] = min(dp[max(0, i-7)]+costs[1], dp[i])
		dp[i] = min(dp[max(0, i-30)]+costs[2], dp[i])
	}

	return dp[len(dp)-1]
}

