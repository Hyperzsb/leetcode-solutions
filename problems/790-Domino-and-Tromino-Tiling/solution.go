func numTilings(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	dp := make([][]int, 2)
	dp[0] = make([]int, n+1)
	dp[1] = make([]int, n+1)
	dp[0][1] = 1
	dp[0][2] = 2
	dp[1][1] = 0
	dp[1][2] = 1
	for i := 3; i <= n; i++ {
		dp[0][i] = (dp[0][i-1] + dp[0][i-2] + 2*dp[1][i-1]) % (1000000000 + 7)
		dp[1][i] = (dp[0][i-2] + dp[1][i-1]) % (1000000000 + 7)
	}
	return dp[0][n]
}

