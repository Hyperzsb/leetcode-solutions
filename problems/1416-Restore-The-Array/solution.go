func numberOfArrays(s string, k int) int {
	dp := make([]int, len(s))

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '0' {
			continue
		}

		if val, _ := strconv.Atoi(string(s[i:])); val <= k {
			dp[i] = 1
		}

		for j := i + 1; j < len(s); j++ {
			if dp[j] == 0 {
				continue
			}

			if val, _ := strconv.Atoi(string(s[i:j])); val <= k {
				dp[i] = (dp[i] + dp[j]) % (1000000000 + 7)
			} else {
				break
			}
		}
	}

	return dp[0]
}

