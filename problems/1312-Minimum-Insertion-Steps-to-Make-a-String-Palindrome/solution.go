func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minInsertions(s string) int {
	dp := make([][]int, len(s))
	for i := 0; i < len(s); i++ {
		dp[i] = make([]int, len(s))
		dp[i][i] = 0
	}

	for i := len(s) - 1; i >= 0; i-- {
		for j := i + 1; j < len(s); j++ {
			if s[i] == s[j] {
				dp[i][j] = dp[i+1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i+1][j]) + 1
			}
		}
	}

	return dp[0][len(s)-1]
}

