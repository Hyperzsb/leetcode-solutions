func max(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func mostPoints(questions [][]int) int64 {
	dp := make([]int64, len(questions))
	dp[len(questions)-1] = int64(questions[len(questions)-1][0])

	for i := len(questions) - 2; i >= 0; i-- {
		if i+questions[i][1]+1 < len(questions) {
			dp[i] = int64(questions[i][0]) + dp[i+questions[i][1]+1]
		} else {
			dp[i] = max(int64(questions[i][0]), dp[i+1])
		}

		dp[i] = max(dp[i], dp[i+1])
	}

	return dp[0]
}

