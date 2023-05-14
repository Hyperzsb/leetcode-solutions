func countGoodStrings(low int, high int, zero int, one int) int {
	dp := make([]int, high+1)
	dp[0] = 1

	for i := 1; i <= high; i++ {
		if i-zero >= 0 {
			dp[i] = (dp[i] + dp[i-zero]) % (1e9 + 7)
		}

		if i-one >= 0 {
			dp[i] = (dp[i] + dp[i-one]) % (1e9 + 7)
		}
	}

	result := 0
	for i := low; i <= high; i++ {
		result = (result + dp[i]) % (1e9 + 7)
	}

	return result
}

