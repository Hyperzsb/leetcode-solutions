func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func stoneGameII(piles []int) int {
	sufSum := make([]int, len(piles))
	sufSum[len(piles)-1] = piles[len(piles)-1]
	for i := len(piles) - 2; i >= 0; i-- {
		sufSum[i] = piles[i] + sufSum[i+1]
	}

	dp := make([][]int, len(piles))
	for i := range dp {
		dp[i] = make([]int, len(piles)+1)
	}

	for i := len(piles) - 1; i >= 0; i-- {
		for m := 1; m <= len(piles); m++ {
			if i+m*2 >= len(piles) {
				dp[i][m] = sufSum[i]
				continue
			}

			for x := 1; x <= 2*m; x++ {
				os := dp[i+x][max(x, m)]
				ts := sufSum[i] - os
				dp[i][m] = max(dp[i][m], ts)
			}
		}
	}

	return dp[0][1]
}

