func coinChange(coins []int, amount int) int {
	sort.Slice(coins, func(a, b int) bool {
		return coins[a] < coins[b]
	})

	dp := make([]int, amount+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = -1
	}
	dp[0] = 0

	for i := coins[0]; i <= amount; i++ {
		result := -1
		for j := 0; j < len(coins) && coins[j] <= i; j++ {
			if dp[i-coins[j]] != -1 {
				if result == -1 {
					result = dp[i-coins[j]] + 1
				} else {
					if result > dp[i-coins[j]]+1 {
						result = dp[i-coins[j]] + 1
					}
				}
			}
		}
		dp[i] = result
	}

	return dp[amount]
}

