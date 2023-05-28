func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func stoneGameIII(stoneValue []int) string {
	n := len(stoneValue)
	stoneValue = append(stoneValue, []int{0, 0}...)

	dp := make([]int, n+3)
	for i := n - 1; i >= 0; i-- {
		takeOne := stoneValue[i] - dp[i+1]
		takeTwo := stoneValue[i] + stoneValue[i+1] - dp[i+2]
		takeThree := stoneValue[i] + stoneValue[i+1] + stoneValue[i+2] - dp[i+3]
		dp[i] = max(takeOne, max(takeTwo, takeThree))
	}

	if dp[0] > 0 {
		return "Alice"
	} else if dp[0] < 0 {
		return "Bob"
	} else {
		return "Tie"
	}
}

