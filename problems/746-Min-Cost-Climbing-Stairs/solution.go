func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minCostClimbingStairs(cost []int) int {
	dp := make([]int, 2)

	for i := 2; i < len(cost); i++ {
		next := min(dp[0]+cost[i-2], dp[1]+cost[i-1])
		dp[0] = dp[1]
		dp[1] = next
	}

	return min(dp[0]+cost[len(cost)-2], dp[1]+cost[len(cost)-1])
}

